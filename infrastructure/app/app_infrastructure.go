package infraApp

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"gitlab.com/kongrentian-group/tianyi/v1/api/controller"
	"gitlab.com/kongrentian-group/tianyi/v1/docs"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"

	infraConfig "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/config"
	infraJWT "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/jwt"
	infraProject "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/project"
	"gitlab.com/kongrentian-group/tianyi/v1/infrastructure/ui/web2"

	infraBranch "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/project/branch"
	infraUser "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/user"

	usecaseApp "gitlab.com/kongrentian-group/tianyi/v1/usecase/app"
	usecaseLifecycle "gitlab.com/kongrentian-group/tianyi/v1/usecase/lifecycle"
	usecaseProject "gitlab.com/kongrentian-group/tianyi/v1/usecase/project"
	usecaseBranch "gitlab.com/kongrentian-group/tianyi/v1/usecase/project/branch"
	usecaseUser "gitlab.com/kongrentian-group/tianyi/v1/usecase/user"
)

var (
	Version     = 1.00
	Title       = "天意 Tianyi"
	Description = "description"
)

type repositories struct {
	user    usecaseUser.Repository
	project usecaseProject.Repository
	branch  usecaseBranch.Repository
}

type server struct {
	config       *infraConfig.App
	database     *gorm.DB
	router       *fiber.App
	controllers  *controller.App
	interactors  *usecaseApp.Interactor
	repositories *repositories
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
	return pkg.NewError(map[string]error{
		"user":    server.repositories.user.Migrate(),
		"project": server.repositories.project.Migrate(),
	})
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
	errors := []error{
		server.router.Shutdown(),
		server.controllers.Session.Storage.Close(),
	}
	for _, err := range errors {
		if err == nil {
			continue
		}
		log.Println("Shutdown error:\n", err)
	}
	if len(errors) != 0 {
		code = 1
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
	server.SetupRepository()
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

func (server *server) SetupRepository() {
	if server.database == nil {
		log.Panicln("database is nil")
	}
	server.repositories = &repositories{
		user:    infraUser.New(server.database),
		project: infraProject.New(server.database),
		branch:  infraBranch.New(server.database),
	}
}

func (server *server) SetupInteractor() {
	if server.repositories == nil {
		log.Panicln("repository is nil")
	}
	jwt := infraJWT.NewInteractor(server.config.Server.JWT)
	branch := usecaseBranch.New(server.repositories.branch)
	server.interactors = &usecaseApp.Interactor{
		Lifecycle: usecaseLifecycle.New(server),
		User:      usecaseUser.New(server.repositories.user, jwt),
		JWT:       jwt,
		Project:   usecaseProject.New(server.repositories.project, branch),
		Branch:    branch,
	}
}

func (server *server) SetupController() {
	if server.repositories == nil {
		log.Panicln("repository is nil")
	}
	server.controllers = &controller.App{
		User: controller.NewUserController(
			server.interactors.User,
		),
		Frontend: controller.NewFrontendController(
			web2.FrontendFilesystem,
			docs.SwaggerFilesystem,
		),
		Session: controller.NewSessionController(
			server.config.Redis,
		),
		Lifecycle: controller.NewLifecycleController(
			server.interactors.Lifecycle,
		),
		Project: controller.NewProjectController(
			server.interactors.Project,
		),
		Branch: controller.NewbranchController(
			server.interactors.Branch,
		),
	}
}

func (server *server) SetupSwagger() {
	docs.SwaggerInfo.Title = Title
	docs.SwaggerInfo.Description = Description
	docs.SwaggerInfo.Version = fmt.Sprintf("%.2f", Version)
	docs.SwaggerInfo.Host = fmt.Sprintf(
		"%s:%s",
		server.config.Server.Host,
		server.config.Server.Port,
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
