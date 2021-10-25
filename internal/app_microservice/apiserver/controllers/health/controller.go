package health

import (
	"github.com/gin-gonic/gin"

	"app_microservice/internal/app_microservice"
	"app_microservice/internal/pkg/service/health"
)

type Controller struct {
	service *health.Service
}

func NewController(service *health.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) RegisterRoutes(r *gin.Engine) {

	apiRoute := r.Group("/api")
	{
		apiRoute.POST("/health", c.Health)
	}
}

func (c *Controller) Health(ctx *gin.Context) {
	ctx.Set(app_microservice.KeyResponse, c.service.Health())
}
