package config

type APIServer struct {
	BaseServer
	Host          string `envconfig:"host"`
	ApiPort       int    `envconfig:"port" default:"8080"`
	TokenPassword string `envconfig:"token_password" default:"password"`
}
