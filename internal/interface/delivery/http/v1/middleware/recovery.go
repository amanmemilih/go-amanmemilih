package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zinct/amanmemilih/config"
	apperr "github.com/zinct/amanmemilih/internal/errors"
	"github.com/zinct/amanmemilih/internal/interface/delivery/http/v1/response"
	"github.com/zinct/amanmemilih/pkg/logger"
)

func Recovery(cfg *config.Config, log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Error("Recovered from panic: %v", r)
				response.JSONError(c, cfg, log, apperr.NewInternalError(fmt.Sprintf("%v", r), nil))
			}
		}()

		c.Next()
	}
}
