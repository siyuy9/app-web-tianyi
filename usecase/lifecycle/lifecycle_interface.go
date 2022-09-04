package usecaseLifecycle

// general lifecycle actions
type Interactor interface {
	// setup everything (but do not run)
	Setup()
	// run the server
	Run()
	// migrate everything
	Migrate() error
}

// specific server actions
type Server interface {
	// listen for requests, blocks the goroutine
	Listen()
	// gracefully shutdown
	Shutdown(exitCode int)
	// gracefully shutdown on panic
	// if there is no panic, do nothing
	ShutdownOnPanic()
	// gracefully shutdown when an interruption signal is sent
	// blocks the goroutine untill the signal is recieved
	// listens for os.Interrupt and syscall.SIGTERM
	ShutdownOnInterruptionSignal()
	// automigrate all repositories
	Migrate() error

	// setup everything
	Setup()
	// populate the config file
	SetupConfig()
	// connect to the database
	SetupDatabase()
	// setup swagger
	SetupSwagger()
	// setup router
	SetupRouter()
	// setup infrastructure
	SetupInfrastructure()
	// setup interactors
	SetupInteractor()
	// setup controllers
	SetupController()
}
