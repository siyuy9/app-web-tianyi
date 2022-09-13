package infra

import (
	infraConfig "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/config"
	infraLifecycle "gitlab.com/kongrentian-group/tianyi/v1/infrastructure/lifecycle"
	useLifecycle "gitlab.com/kongrentian-group/tianyi/v1/usecase/lifecycle"
)

func NewApp(configs ...*infraConfig.App) useLifecycle.Interactor {
	return useLifecycle.New(infraLifecycle.New(configs...).Setup())
}
