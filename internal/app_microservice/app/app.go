package app

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"app_microservice/internal/app_microservice"
	"app_microservice/internal/app_microservice/apiserver"
	"app_microservice/internal/app_microservice/repository"
	"app_microservice/internal/app_microservice/service"
	"app_microservice/internal/app_microservice/storage/postgres"
)

const Name = "app_microservice"

type fxLogger struct {
	logger *zap.SugaredLogger
}

func (fxl fxLogger) Printf(format string, v ...interface{}) {
	fxl.logger.Infof(format, v...)
}

func Provide(conf *app_microservice.Config, zl *zap.Logger) fx.Option {

	return fx.Options(
		fx.StartTimeout(conf.StartTimeout),
		fx.StopTimeout(conf.StopTimeout),

		fx.Logger(
			fxLogger{logger: zl.Named(Name).Sugar()},
		),

		fx.Provide(
			func() *zap.Logger {
				return zl
			},
		),

		fx.Provide(
			func() *app_microservice.Config {
				return conf
			}),

		postgres.Module(),
		repository.Module(),
		service.Module(),
		apiserver.Module(),

		fx.Invoke(
			func(cfg *app_microservice.Config, logger *zap.Logger) {
				logger.Info("Order Management System has started...")
			},
		),
	)
}

func Recover(zl *zap.Logger) {
	if err := recover(); err != nil {
		zl.Fatal("app_microservice recover error", zap.Any("recoveryError", err))
	}
}
