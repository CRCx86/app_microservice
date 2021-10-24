package postgres

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"app_microservice/internal/app_microservice"
	"app_microservice/internal/pkg/storage/postgres"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(func(conf *app_microservice.Config, log *zap.Logger) *postgres.Postgres {
			return postgres.NewPostgres(conf.Postgres, log)
		}),
		fx.Invoke(func(lc fx.Lifecycle, cfg *app_microservice.Config, storage *postgres.Postgres) {
			lc.Append(fx.Hook{
				OnStart: storage.Start,
				OnStop:  storage.Stop,
			})
		}),
	)
}
