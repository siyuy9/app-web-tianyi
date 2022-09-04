package infraPool

import (
	"fmt"

	"github.com/gocraft/work"
)

type Context struct{}

func (context *Context) Log(
	job *work.Job, next work.NextMiddlewareFunc,
) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

func (context *Context) PipelineJob(job *work.Job) error {
	if err := job.ArgError(); err != nil {
		return err
	}
	return nil
}
