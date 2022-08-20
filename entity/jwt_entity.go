package entity

type JWTClaims struct {
	ID    uint `mapstructure:"id" validate:"required,min=1"`
	Admin bool `mapstructure:"admin"`
}
