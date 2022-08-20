package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"gitlab.com/kongrentian-group/tianyi/v1/pkg"
)

// JSON Web Token (JWT) auth middleware.
// For valid token, it sets the user in Ctx.Locals and calls next handler.
// For invalid token, it returns "401 - Unauthorized" error.
// For missing token, it returns "400 - Bad Request" error.
// https://pkg.go.dev/github.com/gofiber/jwt/v3@v3.2.14#section-readme
func NewJWTMiddleware(secret []byte) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: secret,
		// allow registration and login without a token
		Filter: func(context *fiber.Ctx) bool {
			if context.Method() != http.MethodPost {
				return false
			}
			switch context.Path() {
			case "/api/v1/users":
			case "/api/v1/users/login":
			default:
				return false
			}
			return true
		},
		ContextKey:  "user",
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		ErrorHandler: func(context *fiber.Ctx, err error) error {
			// it is written like this in the default handler,
			// couldn't be bothered enough to find actual errors that are being
			// thrown
			if err.Error() == "Missing or malformed JWT" {
				return pkg.NewErrorBadRequest(err)
			}
			return pkg.NewErrorUnauthorized(err)
		},
	})
}
