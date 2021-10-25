package repository

import (
	"app_microservice/internal/pkg/repository/user"
	"go.uber.org/fx"

	"app_microservice/internal/pkg/repository/root"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(root.NewRepository),
		fx.Provide(user.NewRepository),
	)
}
