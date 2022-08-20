package usecaseLifecycle

type interactor struct {
	server Server
}

func New(server Server) Interactor {
	return &interactor{server: server}
}

func (interactor *interactor) Setup() {
	interactor.server.Setup()
}

func (interactor *interactor) Run() {
	go interactor.server.Listen()
	interactor.server.ShutdownOnInterruptionSignal()
}

func (interactor *interactor) Migrate() error {
	return interactor.server.Migrate()
}
