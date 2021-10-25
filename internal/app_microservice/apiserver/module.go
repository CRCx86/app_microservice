package apiserver

import (
	"app_microservice/internal/app_microservice"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"app_microservice/internal/app_microservice/apiserver/controllers/authentication"
	"app_microservice/internal/app_microservice/apiserver/controllers/health"
)

type ApiServer struct {
	fx.In
	Cfg *app_microservice.Config
	Zl  *zap.Logger

	Health *health.Controller
	Auth   *authentication.Controller
}

func Module() fx.Option {
	return fx.Options(

		fx.Provide(health.NewController),
		fx.Provide(authentication.NewController),

		fx.Provide(func(a ApiServer) *APIServer {
			return NewAPIServer(&a.Cfg.APIServer, a.Cfg, a.Zl).
				AddController(a.Health).
				AddController(a.Auth)
		}),

		fx.Invoke(
			func(lf fx.Lifecycle, server *APIServer) {
				lf.Append(fx.Hook{
					OnStart: server.Start,
					OnStop:  server.Stop,
				})
			},
		),
	)
}
