package controller

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

type SessionController *session.Store

func NewSessionController(config *redis.Config) SessionController {
	return session.New(session.Config{
		Storage: redis.New(*config),
	})
}
