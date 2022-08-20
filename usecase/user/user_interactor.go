package usecaseUser

import (
	"github.com/google/uuid"
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
	usecaseJWT "gitlab.com/kongrentian-group/tianyi/v1/usecase/jwt"
)

type userInteractor struct {
	repository Repository
	jwt        usecaseJWT.Interactor
}

func New(repository Repository, jwt usecaseJWT.Interactor) Interactor {
	return &userInteractor{
		repository: repository,
		jwt:        jwt,
	}
}

func (interactor *userInteractor) JWT() usecaseJWT.Interactor {
	return interactor.jwt
}

func (interactor *userInteractor) Validate(user *entity.User) error {
	return pkg.ValidateStruct(user)
}

func (interactor *userInteractor) Create(user *entity.User) error {
	if err := pkg.ValidateStruct(user); err != nil {
		return err
	}
	err := pkg.ValidateStruct(&entity.Password{Value: user.Password})
	if err != nil {
		return err
	}
	user.Password = interactor.PasswordHashCreate(user.Password)
	return interactor.repository.Create(user)
}

func (interactor *userInteractor) Update(user *entity.User) error {
	err := interactor.Validate(user)
	if err != nil {
		return err
	}
	return interactor.repository.Save(user)
}

func (interactor *userInteractor) GetAll() ([]entity.User, error) {
	users, err := interactor.repository.GetAll()
	return users, err
}

func (interactor *userInteractor) Get(id uuid.UUID) (*entity.User, error) {
	return interactor.repository.Get(id)
}

func (interactor *userInteractor) FindByUsername(username string) (
	*entity.User, error,
) {
	return interactor.repository.FindOne(&entity.User{Username: username})
}

func (interactor *userInteractor) PasswordHashCheck(
	user *entity.User, password string,
) (bool, error) {
	return pkg.EncodedHashCompare(password, user.Password)
}

func (interactor *userInteractor) PasswordHashCreate(
	password string,
) string {
	return pkg.EncodedHashGenerate(password)
}
