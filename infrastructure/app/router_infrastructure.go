package infraApp

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gitlab.com/kongrentian-group/tianyi/v1/api/controller"
	"gitlab.com/kongrentian-group/tianyi/v1/api/middleware"
	"gitlab.com/kongrentian-group/tianyi/v1/api/presenter"
	_ "gitlab.com/kongrentian-group/tianyi/v1/docs"
	infraConfig "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/config"
)

func newRouter(config *fiber.Config) *fiber.App {
	config.ErrorHandler = presenter.Error
	return fiber.New(*config)
}

func setupRouter(
	router *fiber.App,
	app *controller.App,
	config *infraConfig.App,
) {
	// use global middlewares
	// https://docs.gofiber.io/api/middleware
	router.Use(
		// limiter.New(),
		cors.New(),
		compress.New(),
		etag.New(),
		// csrf.New(csrf.Config{CookieSameSite: "None"}),
		logger.New(),
		recover.New(recover.Config{EnableStackTrace: true}),
	)

	// redirect for gitlab-runners

	// groups
	apiRoot := router.Group(
		"api",
		middleware.NewJWTMiddleware(config.Server.JWT.GetSecret()),
	)
	api := apiRoot.Group("v1")
	users := api.Group("users")
	projects := api.Group("projects")
	database := api.Group("database")
	project := projects.Group(":project_id")
	branches := project.Group("branches")
	branch := branches.Group(":branch_name")
	pipelines := branch.Group("pipelines")
	pipeline := pipelines.Group(":pipeline_name")

	// user routes
	users.Get("", app.User.GetAll)
	users.Post("", app.User.Create)
	users.Post("login", app.User.Login)
	users.Get("user/:user_id", app.User.Get)

	// project routes
	projects.Get("", app.Project.Get)
	projects.Post("", app.Project.Create)

	project.Get("", app.Project.GetByID)
	project.Put("", app.Project.Update)

	branches.Post("", app.Branch.Create)
	branches.Get("", app.Branch.GetProjectBranches)

	branch.Get("", app.Branch.Get)
	branch.Put("", app.Branch.Update)

	pipeline.Post("", app.Pipeline.Create)

	// swagger routes
	api.Use("swagger", app.Frontend.ServeSwagger)

	// database routes
	database.Post("migrate", app.Lifecycle.Migrate)

	// api catch all in order to not send html in response to
	// invalid api requests
	apiRoot.Use(presenter.RouteDoesNotExist)

	// frontend controller
	//
	// returns embedded files from ../infrastructure/ui/web2/dist
	// (the folder is created and populated by 'yarn build')
	//
	// it is a Single Page Application, routing is done by Javascript, that's
	// why it has to return index.html on missing files
	router.Use(app.Frontend.Serve)
}
