package controller

type App struct {
	User      interface{ UserController }
	Frontend  FrontendController
	Session   SessionController
	Lifecycle LifecycleController
	Project   ProjectController
	Branch    BranchController
	Pipeline  PipelineController
}
