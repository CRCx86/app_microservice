package health

import (
	"app_microservice/internal/app_microservice"
)

type Service struct {
	cfg *app_microservice.Config
}

func NewService(cfg *app_microservice.Config) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) Health() map[string]string {
	return map[string]string{"status": "ok"}
}
