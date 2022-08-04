package user

import (
	"errors"

	"gorm.io/gorm"

	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model/db"
)

// took some stuff from
// https://github.com/gothinkster/golang-gin-realworld-example-app/blob/master/users/models.go
type Store interface {
	// return a model based on the condition
	// 	userModel, err := FindOneUser(&UserModel{Username: "username0"})
	Find(condition interface{}) (*db.UserModel, error)

	// save update value in database
	// if the value doesn't have a primary key, it will be inserted
	// 	err := SaveOne(&userModel)
	Save(data interface{}) error
}

type store struct {
	database *gorm.DB
}

func NewStore(database *gorm.DB) Store {
	return &store{database: database}
}

func (store *store) Find(
	condition interface{},
) (*db.UserModel, error) {
	var model *db.UserModel
	return model, store.database.Where(condition).First(model).Error
}

func (store *store) Save(data interface{}) error {
	return store.database.Save(data).Error
}

// set password (hashed)
func (store *store) SetPassword(
	user *db.UserModel, password string,
) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	encodedHash, err := generateEncodedHash(password)
	if err == nil {
		user.PasswordHash = encodedHash
	}
	return err
}

// check if string is equal to the hashed password of the user
func (store *store) CheckPassword(
	user *db.UserModel, password string,
) (match bool, err error) {
	return comparePasswordAndEncodedHash(password, user.PasswordHash)
}
