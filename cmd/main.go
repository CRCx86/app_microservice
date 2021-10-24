package main

import (
	"go.uber.org/fx"

	"app_microservice/internal/app_microservice"
	"app_microservice/internal/app_microservice/app"
	"app_microservice/internal/pkg/logger"
	"app_microservice/internal/pkg/tracing"
)

var (
	version   string
	buildDate string
	commit    string
)

func main() {

	conf, err := app_microservice.NewConfig()
	if err != nil {
		panic(err)
	}

	conf.Version = version
	conf.BuildDate = buildDate
	conf.Commit = commit

	zapLogger, err := logger.New(app.Name, *conf)
	if err != nil {
		panic(err)
	}
	tracing.New(zapLogger)

	defer app.Recover(zapLogger)
	fx.New(
		app.Provide(conf, zapLogger),
	).Run()

}
