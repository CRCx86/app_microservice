package log

import (
	"context"
	"go.uber.org/zap"

	"app_microservice/internal/app_microservice"
	v7 "app_microservice/internal/pkg/storage/elastic/v7"
	"app_microservice/internal/pkg/util"
)

type Service struct {
	zl      *zap.Logger
	cfg     *app_microservice.Config
	storage *v7.Elastic
}

func NewService(cfg *app_microservice.Config, storage *v7.Elastic, zl *zap.Logger) *Service {
	return &Service{
		zl:      zl,
		cfg:     cfg,
		storage: storage,
	}
}

func (s *Service) LogMessage(
	c context.Context,
	messageType string,
	message string,
	indexElastic string,
	data map[string]interface{}) {

	go func() {
		_, err := s.storage.Create(c,
			util.MessageToExternalLog(
				data,
				messageType,
				message),
			"",
			indexElastic)
		s.zl.Sugar().Info(message)
		if err != nil {
			s.zl.Sugar().Info(err)
		}
	}()

}
