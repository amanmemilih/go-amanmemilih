package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
}

func JSONSuccess(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code:    http.StatusOK,
		Message: message,
		Success: true,
		Data:    data,
	})
}
