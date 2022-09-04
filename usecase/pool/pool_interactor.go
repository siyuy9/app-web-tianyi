package usecasePool

type interactor struct {
	pool Pool
}

func New(pool Pool) Interactor {
	return &interactor{pool}
}
