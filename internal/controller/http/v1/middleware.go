package v1

import (
	"db_cp_6/internal/service"
	"db_cp_6/pkg/logger"
	"github.com/gin-gonic/gin"
	pkgErrors "github.com/pkg/errors"
	"net/http"
)

type AuthMiddleware struct {
	authService service.Auth
	log         *logger.Logger
}

func (m *AuthMiddleware) SessionCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query("token")

		if !m.authService.GetSession(token) {
			m.log.Errorf("AuthMiddleware SessionCheck: %v", service.ErrSessionNotExists)
			ctx.JSON(http.StatusBadRequest, pkgErrors.WithMessage(service.ErrSessionNotExists, token))
			return
		}

		ctx.Next()
	}
}
