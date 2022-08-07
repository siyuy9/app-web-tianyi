package interactor

import (
	"sync"

	validatorPackage "github.com/go-playground/validator/v10"

	"gitlab.com/kongrentian-group/tianyi/entity"
	"gitlab.com/kongrentian-group/tianyi/infrastructure/registry"
	"gitlab.com/kongrentian-group/tianyi/pkg"
)

var (
	userInteractorInstance UserInteractor
	once                   sync.Once
)

// UserService : represent the user's services
type UserInteractor interface {
	// check whether the given user is valid
	Validate(user *entity.User) error
	// return all users in the database
	FindAll() ([]entity.User, error)
	// check whether a given unhashed password matches the hashed password of
	// the user
	PasswordCheck(user *entity.User, password string) (matches bool, err error)
	// creates a hash from the given string
	PasswordCreate(password string) (hashedPassword string, err error)
}

type userInteractor struct {
	registry  registry.UserRegistry
	validator *validatorPackage.Validate
}

// NewuserInteractor: construction function, injected by user repository
func NewUserInteractor(registry registry.UserRegistry) UserInteractor {
	once.Do(func() {
		userInteractorInstance = &userInteractor{
			registry:  registry,
			validator: validatorPackage.New(),
		}
	})
	return userInteractorInstance
}

func (service *userInteractor) Validate(user *entity.User) error {
	return service.validator.Struct(user)
}

func (service *userInteractor) FindAll() ([]entity.User, error) {
	users, err := service.registry.FindAll()
	return users, err
}

func (service *userInteractor) PasswordCheck(
	user *entity.User, password string,
) (bool, error) {
	return pkg.EncodedHashCompare(password, *user.PasswordHash)
}

// TODO: add some password
func (service *userInteractor) PasswordCreate(password string) (string, error) {
	return pkg.EncodedHashGenerate(password)
}
