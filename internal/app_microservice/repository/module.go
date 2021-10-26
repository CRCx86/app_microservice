package repository

import (
	"go.uber.org/fx"

	"app_microservice/internal/pkg/repository/root"
	"app_microservice/internal/pkg/repository/user"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(root.NewRepository),
		fx.Provide(user.NewRepository),
	)
}
