package main

import (
	"go.uber.org/fx"
	"os"

	"app_microservice/internal/app_microservice"
	"app_microservice/internal/app_microservice/app"
	"app_microservice/internal/pkg/tracing"
	"app_microservice/internal/pkg/zaplog"
)

var (
	version   string
	buildDate string
	commit    string
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		app_microservice.Usage()
		return
	}
	conf, err := app_microservice.NewConfig()
	if err != nil {
		panic(err)
	}

	conf.Version = version
	conf.BuildDate = buildDate
	conf.Commit = commit

	zapLogger, err := zaplog.New(app.Name, *conf)
	if err != nil {
		panic(err)
	}
	tracing.New(zapLogger)

	defer app.Recover(zapLogger)
	fx.New(
		app.Provide(conf, zapLogger),
	).Run()

}
