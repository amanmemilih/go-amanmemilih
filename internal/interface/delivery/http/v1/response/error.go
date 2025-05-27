package response

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/zinct/amanmemilih/config"
	apperr "github.com/zinct/amanmemilih/internal/errors"
	"github.com/zinct/amanmemilih/pkg/logger"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Stack   string `json:"stack,omitempty"`
}

func JSONError(c *gin.Context, cfg *config.Config, log *logger.Logger, err error) {
	appErr, ok := err.(*apperr.APPError)
	if !ok {
		appErr = apperr.NewInternalError(err.Error(), nil)
	}

	if cfg.App.Env == "production" {
		log.Error(err.Error())

		c.JSON(appErr.Code, ErrorResponse{
			Code:    appErr.Code,
			Message: appErr.Message,
			Success: false,
			Data:    appErr.Data,
		})
		return
	} else {
		log.Error(err.Error() + " - " + string(debug.Stack()))

		c.JSON(appErr.Code, ErrorResponse{
			Code:    appErr.Code,
			Message: appErr.Message,
			Success: false,
			Data:    appErr.Data,
			Stack:   string(debug.Stack()),
		})
	}
}
