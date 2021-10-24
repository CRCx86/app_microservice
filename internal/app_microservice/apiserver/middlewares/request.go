package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"

	"app_microservice/internal/app_microservice"
)

func RequestMiddleware(zl *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request

		c.Set("requestId", uuid.NewV4().String())

		var req app_microservice.Request

		if c.Request.Method == http.MethodPost {
			err := c.ShouldBindJSON(&req)
			if err != nil {
				zl.Sugar().Error(c, err.Error())

				c.JSON(http.StatusOK, map[string]string{"error": err.Error()})

				c.Abort()

				return
			}

			c.Set(app_microservice.KeyMeta, req.Meta)
			c.Set(app_microservice.KeyRequest, req.Data)
			zl.Sugar().Info(c, string(req.Meta))
			zl.Sugar().Info(c, string(req.Data))
		}

		c.Next()
	}
}
