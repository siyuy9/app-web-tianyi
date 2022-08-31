package usecasePipeline

type interactor struct{}

func New() Interactor {
	return &interactor{}
}
