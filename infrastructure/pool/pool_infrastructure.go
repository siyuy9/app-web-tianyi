package infraPool

import (
	"fmt"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	infraConfig "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/config"

	usecasePool "gitlab.com/kongrentian-group/tianyi/v1/usecase/pool"
)

type pool struct {
	workerPool *work.WorkerPool
	redisPool  *redis.Pool
	enqueuer   *work.Enqueuer
}

func New(config *infraConfig.App) usecasePool.Pool {
	redisPool := &redis.Pool{
		MaxActive: 15,
		MaxIdle:   15,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
			)
		},
	}
	namespace := "{tianyi}"
	workerPool := work.NewWorkerPool(Context{}, 10, namespace, redisPool)
	enqueuer := work.NewEnqueuer(namespace, redisPool)
	workerPool.Middleware((*Context).Log)

	workerPool.Job("pipeline_job", (*Context).PipelineJob)

	return &pool{workerPool, redisPool, enqueuer}
}

func (pool *pool) Start() {
	pool.workerPool.Start()
}

func (pool *pool) Close() error {
	pool.workerPool.Stop()
	return pool.redisPool.Close()
}
