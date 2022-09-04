package usecasePool

type Interactor interface{}

type Pool interface {
	// start the pool
	Start()
	// close the pool
	Close() error
}
