package infraPool

import (
	"fmt"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	infraConfig "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/config"
	usecaseJob "gitlab.com/kongrentian-group/tianyi/v1/usecase/job"
	usecasePool "gitlab.com/kongrentian-group/tianyi/v1/usecase/pool"
)

type pool struct {
	workerPool    *work.WorkerPool
	redisPool     *redis.Pool
	enqueuer      *work.Enqueuer
	jobRepository usecaseJob.Repository
}

type workJob func(*work.Job) error

func New(
	config *infraConfig.App, jobRepository usecaseJob.Repository,
) usecasePool.Pool {
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
	return &pool{workerPool, redisPool, enqueuer, jobRepository}
}

func (pool *pool) Start() {
	pool.workerPool.Start()
}

func (pool *pool) Close() error {
	pool.workerPool.Stop()
	return pool.redisPool.Close()
}

func (pool *pool) Job(
	name string, handler usecasePool.Handler,
	errorHandler usecasePool.ErrorHandler,
) {
	pool.workerPool.JobWithOptions(
		name, work.JobOptions{MaxFails: 1},
		pool.normalizeHandler(handler, errorHandler),
	)
}

func (pool *pool) normalizeHandler(
	handler usecasePool.Handler, errorHandler usecasePool.ErrorHandler,
) workJob {
	return func(jobRedis *work.Job) error {
		job, err := pool.jobRepository.GetByRedisID(jobRedis.ID)
		if err != nil {
			return fmt.Errorf(
				"could not get the job from the database: %w", err,
			)
		}
		if err = handler(job); err != nil {
			return errorHandler(job, err)
		}
		return nil
	}
}
