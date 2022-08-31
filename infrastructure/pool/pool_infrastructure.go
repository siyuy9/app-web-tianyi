package infraPipeline

import (
	"github.com/gomodule/redigo/redis"
)

type infrastructure struct{}

func New(pool *redis.Pool) *infrastructure {
	return nil
}
