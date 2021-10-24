package repository

import (
	"go.uber.org/fx"

	"app_microservice/internal/pkg/repository/root"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(root.NewRepository),
	)
}
