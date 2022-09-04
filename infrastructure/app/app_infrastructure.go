package infraApp

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"gitlab.com/kongrentian-group/tianyi/v1/api/controller"
	"gitlab.com/kongrentian-group/tianyi/v1/docs"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
	web2 "gitlab.com/kongrentian-group/tianyi/v1/web"

	infraBranch "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/branch"
	infraConfig "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/config"
	infraJob "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/job"
	infraJWT "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/jwt"
	infraPipeline "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/pipeline"
	infraPool "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/pool"
	infraProject "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/project"
	infraSession "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/session"
	infraUser "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/user"

	usecaseApp "gitlab.com/kongrentian-group/tianyi/v1/usecase/app"
	usecaseBranch "gitlab.com/kongrentian-group/tianyi/v1/usecase/branch"
	usecaseJob "gitlab.com/kongrentian-group/tianyi/v1/usecase/job"
	usecaseLifecycle "gitlab.com/kongrentian-group/tianyi/v1/usecase/lifecycle"
	usecasePipeline "gitlab.com/kongrentian-group/tianyi/v1/usecase/pipeline"
	usecasePool "gitlab.com/kongrentian-group/tianyi/v1/usecase/pool"
	usecaseProject "gitlab.com/kongrentian-group/tianyi/v1/usecase/project"

	usecaseSession "gitlab.com/kongrentian-group/tianyi/v1/usecase/session"
	usecaseUser "gitlab.com/kongrentian-group/tianyi/v1/usecase/user"
)

var (
	Version     = 1.00
	Title       = "天意 Tianyi"
	Description = "description"
)

type infrastructure struct {
	user     usecaseUser.Repository     `validate:"required"`
	project  usecaseProject.Repository  `validate:"required"`
	branch   usecaseBranch.Repository   `validate:"required"`
	session  usecaseSession.Interactor  `validate:"required"`
	pool     usecasePool.Pool           `validate:"required"`
	job      usecaseJob.Repository      `validate:"required"`
	pipeline usecasePipeline.Repository `validate:"required"`
}

type server struct {
	config         *infraConfig.App
	database       *gorm.DB
	router         *fiber.App
	controllers    *controller.App
	interactors    *usecaseApp.Interactor
	infrastructure *infrastructure
}

func new(configs ...*infraConfig.App) *server {
	var config *infraConfig.App
	if len(configs) != 0 {
		config = configs[0]
	} else {
		config = infraConfig.New()
	}

	return &server{config: config}
}

func New(configs ...*infraConfig.App) *usecaseApp.Interactor {
	server := new(configs...)
	server.Setup()
	return server.interactors
}

func (server *server) Migrate() error {
	repositories := []interface{ Migrate() error }{
		server.infrastructure.branch,
		server.infrastructure.project,
		server.infrastructure.user,
	}
	for _, repository := range repositories {
		if err := repository.Migrate(); err != nil {
			return err
		}
	}
	return nil
}

func (server *server) Run() {
	go server.Listen()
	server.ShutdownOnInterruptionSignal()
}

func (server *server) Listen() {
	defer server.ShutdownOnPanic()
	err := server.router.Listen(":" + server.config.Server.Port)
	if err == nil {
		return
	}
	log.Println("Runtime error:\n", err)
	server.Shutdown(1)
}

func (server *server) ShutdownOnPanic() {
	err := recover()
	if err == nil {
		return
	}
	log.Println("Runtime panic:\n", err)
	server.Shutdown(1)
}

func (server *server) Shutdown(code int) {
	log.Println("Shutting down...")
	shutdowns := []func() error{
		server.router.Shutdown,
		server.infrastructure.session.Close,
		server.infrastructure.pool.Close,
	}
	waitGroup := sync.WaitGroup{}
	failToggle := false
	for _, shutdown := range shutdowns {
		waitGroup.Add(1)
		go func(shutdown func() error) {
			defer waitGroup.Done()
			if err := shutdown(); err != nil {
				failToggle = true
				log.Println(err)
			}
		}(shutdown)
	}
	waitGroup.Wait()
	if failToggle {
		os.Exit(1)
	}
	os.Exit(code)
}

func (server *server) ShutdownOnInterruptionSignal() {
	// channel for the interruption signal
	channel := make(chan os.Signal, 1)
	// when an interruption or termination signal is sent, shutdown
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	// wait for it and shutdown
	<-channel
	server.Shutdown(0)
}

func (server *server) Setup() {
	server.SetupConfig()
	server.SetupDatabase()
	server.SetupInfrastructure()
	server.SetupInteractor()
	server.SetupController()
	server.SetupSwagger()
	server.SetupRouter()
}

func (server *server) SetupConfig() {
	server.config.Populate()
}

func (server *server) SetupDatabase() {
	server.database = connectDatabase(server.config.Database)
}

func (server *server) SetupInfrastructure() {
	if server.database == nil {
		log.Panicln("database is nil")
	}
	job := infraJob.New(server.database)
	server.infrastructure = &infrastructure{
		user:     infraUser.New(server.database),
		project:  infraProject.New(server.database),
		branch:   infraBranch.New(server.database),
		session:  infraSession.New(server.config.Redis),
		pool:     infraPool.New(server.config, job),
		job:      job,
		pipeline: infraPipeline.New(server.database),
	}
	if err := pkg.ValidateStruct(server.infrastructure); err != nil {
		panic(err)
	}
}

func (server *server) SetupInteractor() {
	if server.infrastructure == nil {
		log.Panicln("repository is nil")
	}
	branch := usecaseBranch.New(server.infrastructure.branch)
	job := usecaseJob.New(server.infrastructure.job)
	server.interactors = &usecaseApp.Interactor{
		Lifecycle: usecaseLifecycle.New(server),
		User:      usecaseUser.New(server.infrastructure.user),
		JWT:       infraJWT.NewInteractor(server.config.Server.JWT),
		Project:   usecaseProject.New(server.infrastructure.project, branch),
		Branch:    branch,
		Session:   usecaseSession.New(server.infrastructure.session),
		Pool:      usecasePool.New(server.infrastructure.pool),
		Job:       job,
		Pipeline:  usecasePipeline.New(server.infrastructure.pipeline, job),
	}
	if err := pkg.ValidateStruct(server.interactors); err != nil {
		log.Panicln(err)
	}
}

func (server *server) SetupController() {
	if server.infrastructure == nil {
		log.Panicln("repository is nil")
	}
	jwt := controller.NewJWT(server.config.Server.JWT.GetSecret())
	server.controllers = &controller.App{
		User: controller.NewUser(
			server.interactors.User, server.interactors.Session,
		),
		Frontend: controller.NewFrontend(
			web2.FrontendFilesystem, docs.SwaggerFilesystem,
		),
		Session: controller.NewSession(
			server.infrastructure.session, jwt, server.interactors.JWT,
		),
		Lifecycle: controller.NewLifecycle(
			server.interactors.Lifecycle,
		),
		Project: controller.NewProject(
			server.interactors.Project,
		),
		Branch: controller.NewBranch(
			server.interactors.Branch, server.interactors.Project,
		),
		Pipeline: controller.NewPipeline(
			server.interactors.Branch, server.interactors.Project,
		),
		JWT: jwt,
	}
	if err := pkg.ValidateStruct(server.controllers); err != nil {
		log.Panicln(err)
	}
	server.infrastructure.pool.Job(
		usecasePipeline.InteractorKey, server.interactors.Pipeline.RunJob,
		server.interactors.Pipeline.JobErrorHandler,
	)
}

func (server *server) SetupSwagger() {
	docs.SwaggerInfo.Title = Title
	docs.SwaggerInfo.Description = Description
	docs.SwaggerInfo.Version = fmt.Sprintf("%.2f", Version)
	docs.SwaggerInfo.Host = fmt.Sprintf(
		"%s:%s", server.config.Server.Host, server.config.Server.Port,
	)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func (server *server) SetupRouter() {
	if server.controllers == nil {
		log.Panicln("controller is nil")
	}
	server.router = newRouter(server.config.Fiber)
	setupRouter(server.router, server.controllers, server.config)
}
