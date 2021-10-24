package middlewares

import (
	"encoding/json"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"

	"app_microservice/internal/app_microservice"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// after request

		if c.Request.Method == http.MethodPost {
			meta := c.MustGet(app_microservice.KeyMeta).(json.RawMessage)
			result, _ := json.Marshal(c.MustGet(app_microservice.KeyResponse))
			var response interface{}

			isError := len(c.Errors) > 0
			if isError {
				stack := strings.Split(string(debug.Stack()), "\n")
				response = app_microservice.ResponseError{
					Success:  0,
					Envelope: app_microservice.Envelope{Meta: meta},
					Error: app_microservice.RError{
						Message:    result,
						StackTrace: stack,
					},
				}
			} else {
				response = app_microservice.ResponseSuccess{
					Success:  1,
					Envelope: app_microservice.Envelope{Meta: meta},
					Data:     result,
				}
			}

			c.JSON(http.StatusOK, response)
		}
	}
}
