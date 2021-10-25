package middlewares

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"app_microservice/internal/app_microservice"
	"app_microservice/internal/pkg/service/user"
)

func JwtAuthenticationMiddleware(zl *zap.Logger, cfg *app_microservice.Config) gin.HandlerFunc {
	return JwtAuthentication(zl, cfg)
}

func JwtAuthentication(l *zap.Logger, cfg *app_microservice.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		noAuth := []string{"/api/account/auth", "/api/account/login"}
		requestPath := ctx.Request.URL.Path

		for _, value := range noAuth {
			if value == requestPath {
				ctx.Next()
				return
			}
		}

		tokenHeader := ctx.Request.Header.Get("Authorization")

		if tokenHeader == "" {
			_ = ctx.Error(errors.New("missing auth token"))
			ctx.Set(app_microservice.KeyResponse, "missing auth token")
			ctx.Next()
			return
		}

		splitter := strings.Split(tokenHeader, " ")
		if len(splitter) != 2 {
			_ = ctx.Error(errors.New("invalid/Malformed auth token"))
			ctx.Set(app_microservice.KeyResponse, "invalid/Malformed auth token")
			ctx.Next()
			return
		}

		tokenPart := splitter[1]
		tk := &user.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.APIServer.TokenPassword), nil
		})

		if err != nil {
			_ = ctx.Error(errors.New("malformed user token"))
			ctx.Set(app_microservice.KeyResponse, "malformed auth token")
			ctx.Next()
			return
		}

		if !token.Valid {
			_ = ctx.Error(errors.New("token is not valid"))
			ctx.Set(app_microservice.KeyResponse, "token is not valid")
			ctx.Next()
			return
		}

		ctx.Set(app_microservice.KeyUser, tk.UserId)

		ctx.Next()
	}
}
