package user

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

	apiRoute := r.Group("/api/user")
	{
		apiRoute.POST("/create", c.Create)
		apiRoute.POST("/list", c.List)
		apiRoute.POST("/get", c.Get)
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

func (c *Controller) List(ctx *gin.Context) {

	list, err := c.service.List(ctx)
	if err != nil {
		ctx.Set(app_microservice.KeyResponse, err.Error())
		return
	}

	ctx.Set(app_microservice.KeyResponse, list)
}

func (c *Controller) Get(ctx *gin.Context) {

	data := ctx.Keys[app_microservice.KeyRequest].(json.RawMessage)

	var item dto.User
	err := json.Unmarshal(data, &item)
	if err != nil {
		ctx.Set(app_microservice.KeyResponse, err.Error())
		return
	}

	userDB, err := c.service.GetById(ctx, item)
	if err != nil {
		ctx.Set(app_microservice.KeyResponse, err.Error())
		return
	}

	ctx.Set(app_microservice.KeyResponse, userDB)
}
