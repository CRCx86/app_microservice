package authentication

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"app_microservice/internal/app_microservice"
	"app_microservice/internal/pkg/dto"
	"app_microservice/internal/pkg/service/user"
)

type Controller struct {
	service *user.Service
}

func NewController(service *user.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) RegisterRoutes(r *gin.Engine) {

	apiRoute := r.Group("/api/account")
	{
		apiRoute.POST("/auth", c.Create)
		apiRoute.POST("/login", c.Login)
	}
}

func (c *Controller) Create(ctx *gin.Context) {
	data := ctx.Keys[app_microservice.KeyRequest].(json.RawMessage)

	var item dto.User
	err := json.Unmarshal(data, &item)
	if err != nil {
		ctx.Set(app_microservice.KeyResponse, err.Error())
		return
	}

	token, err := c.service.Create(ctx, item)
	if err != nil {
		ctx.Set(app_microservice.KeyResponse, err.Error())
		return
	}

	ctx.Set(app_microservice.KeyResponse, gin.H{"token: ": token})
}

func (c *Controller) Login(ctx *gin.Context) {
	data := ctx.Keys[app_microservice.KeyRequest].(json.RawMessage)

	var item dto.User
	err := json.Unmarshal(data, &item)
	if err != nil {
		ctx.Set(app_microservice.KeyResponse, err.Error())
		return
	}

	account, err := c.service.Login(ctx, item)
	if err != nil {
		ctx.Set(app_microservice.KeyResponse, err.Error())
		return
	}

	ctx.Set(app_microservice.KeyResponse, account)
}
