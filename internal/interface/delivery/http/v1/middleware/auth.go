package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zinct/amanmemilih/config"
	apperr "github.com/zinct/amanmemilih/internal/errors"
	"github.com/zinct/amanmemilih/internal/interface/delivery/http/v1/response"
	"github.com/zinct/amanmemilih/pkg/jwt"
	"github.com/zinct/amanmemilih/pkg/logger"
)

func JWTAuthMiddleware(jm *jwt.JWTManager, cfg *config.Config, log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.JSONError(c, cfg, log, apperr.NewUnauthorizedError("Authorization header is missing", nil))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.JSONError(c, cfg, log, apperr.NewUnauthorizedError("Authorization header format must be Bearer {token}", nil))
			return
		}

		tokenStr := parts[1]
		claims, err := jm.ValidateJWT(tokenStr)
		if err != nil {
			response.JSONError(c, cfg, log, apperr.NewUnauthorizedError("Invalid or expired token", nil))
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("token", tokenStr)

		c.Next()
	}
}
