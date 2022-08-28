package infraJWT

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/mapstructure"
	"gitlab.com/kongrentian-group/tianyi/v1/entity"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"

	infraConfig "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/config"
	usecaseJWT "gitlab.com/kongrentian-group/tianyi/v1/usecase/jwt"
)

type jwtInteractor struct {
	config *infraConfig.JWT
	secret []byte
}

func NewInteractor(config *infraConfig.JWT) usecaseJWT.Interactor {
	return &jwtInteractor{config: config, secret: config.GetSecret()}
}

func (interactor *jwtInteractor) New(user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"admin": user.Admin,
		"exp": time.Now().Add(
			time.Hour * time.Duration(interactor.config.Expiration),
		).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(interactor.secret)
}

func (interactor *jwtInteractor) GetClaims(token interface{}) (
	*entity.JWTClaims, error,
) {
	if token == nil {
		return nil, fmt.Errorf("JWT is nil")
	}
	tokenAsserted, ok := token.(*jwt.Token)
	if !ok {
		return nil, fmt.Errorf(
			"cannot get claims, invalid token: %v", token,
		)
	}
	claimsMap, ok := tokenAsserted.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf(
			"cannot get claims, invalid claim map: %v", tokenAsserted.Claims,
		)
	}
	claims := &entity.JWTClaims{}
	if err := mapstructure.Decode(token, claimsMap); err != nil {
		return nil, err
	}
	return claims, pkg.ValidateStruct(claims)
}
