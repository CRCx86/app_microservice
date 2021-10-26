package service

import (
	"go.uber.org/fx"

	"app_microservice/internal/pkg/service/health"
	"app_microservice/internal/pkg/service/user"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(health.NewService),
		fx.Provide(user.NewService),

		//fx.Invoke(func(lc fx.Lifecycle, cfg *app_microservice.Config, service *robot.Service) {
		//	lc.Append(fx.Hook{
		//		OnStart: service.Start,
		//		OnStop:  service.Stop,
		//	})
		//}),
	)
}
