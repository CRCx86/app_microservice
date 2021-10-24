package app_microservice

import (
	"encoding/json"
	"time"

	"github.com/kelseyhightower/envconfig"

	"app_microservice/internal/pkg/config"
	"app_microservice/internal/pkg/storage/postgres"
)

const CoreEnvironmentPrefix = "APP_MICROSERVICE"

const EnvDev = "dev"

type Config struct {
	Env                string           `envconfig:"env"`
	Debug              bool             `envconfig:"debug"`
	ProfilerEnable     bool             `envconfig:"pprof"`
	StartTimeout       time.Duration    `envconfig:"start_timeout" default:"20s"`
	StopTimeout        time.Duration    `envconfig:"stop_timeout" default:"60s"`
	APIServer          config.APIServer `envconfig:"apiserver"`
	Storage            config.Storage   `envconfig:"storage"`
	Postgres           postgres.Config  `envconfig:"postgres"`
	Logger             config.Logger    `envconfig:"logger"`
	MaxCollectTime     time.Duration    `envconfig:"max_collect_time" default:"10m"`
	MaxRobotGoroutines int              `envconfig:"max_robot_goroutines" default:"10"`
	Version            string
	BuildDate          string
	Commit             string
}

func NewConfig() (*Config, error) {

	cfg := &Config{}

	if err := envconfig.Process(CoreEnvironmentPrefix, cfg); err != nil {
		return nil, err
	}

	if cfg.Debug {
		cfg.Logger.Level = "debug"
		cfg.Logger.Debug = true
	}

	return cfg, nil
}

const (
	KeyMeta = "meta"

	KeyRequest  = "requestData"
	KeyResponse = "responseData"
)

type Envelope struct {
	Meta json.RawMessage `json:"meta"`
}

type Request struct {
	Envelope
	Data json.RawMessage `json:"data" binding:"required"`
}

type ResponseSuccess struct {
	Success int `json:"success"`
	Envelope
	Data json.RawMessage `json:"data" binding:"required"`
}

type RError struct {
	Message    json.RawMessage `json:"message" binding:"required"`
	StackTrace []string        `json:"stackTrace" binding:"required"`
}

type ResponseError struct {
	Success int `json:"success"`
	Envelope
	Error RError `json:"error" binding:"required"`
}
