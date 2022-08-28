// https://docs.gofiber.io/api/middleware/session
package infraSession

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	usecaseSession "gitlab.com/kongrentian-group/tianyi/v1/usecase/session"
)

type repository struct {
	session *session.Store
}

func New(config *redis.Config) usecaseSession.Repository {
	return &repository{
		session: session.New(session.Config{
			Storage: redis.New(*config),
		}),
	}
}

func (repository *repository) Get(context interface{}) (
	usecaseSession.Session, error,
) {
	contextAsserted, ok := context.(*fiber.Ctx)
	if !ok {
		return nil, fmt.Errorf("invalid context type: %v+", context)
	}
	return repository.session.Get(contextAsserted)
}

func (repository *repository) Reset() error {
	return repository.session.Reset()
}

func (repository *repository) Close() error {
	return repository.session.Storage.Close()
}
