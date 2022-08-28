package usecaseJWT

import "gitlab.com/kongrentian-group/tianyi/v1/entity"

type Interactor interface {
	// generate user jwt
	New(user *entity.User) (string, error)
	// get claims from jwt
	GetClaims(token interface{}) (*entity.JWTClaims, error)
}
