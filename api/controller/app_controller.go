package controller

type AppController struct {
	User      interface{ UserController }
	Frontend  FrontendController
	Lifecycle LifecycleController
}
