package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/kongrentian-groups/golang/tianyi/backend/model/request/auth"
)

// SignUp godoc
// @Summary Register a new user
// @Description Register a new user
// @ID sign-up
// @Tags user
// @Accept json
// @Produce json
// @Param user body auth.LoginRequest true "User info for registration"
// @Success 200 {object} auth.LoginResponse
// @Failure 403 {object} auth.LoginResponse
// @Router /users [post]
func loginHandler(context *fiber.Ctx) error {
	request := &auth.LoginRequest{}
	response := &auth.LoginResponse{}
	log.Println(request, response)
	return nil
}
