package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
	"github.com/zinct/amanmemilih/pkg/logger"
)

type ProvinceController struct {
	uc  usecases.ProvinceUsecase
	cfg *config.Config
	log *logger.Logger
}

func NewProvinceController(provinceUsecase usecases.ProvinceUsecase, cfg *config.Config, log *logger.Logger) *ProvinceController {
	return &ProvinceController{uc: provinceUsecase, cfg: cfg, log: log}
}

func (c *ProvinceController) FindAll(ctx *gin.Context) {
	provinces, err := c.uc.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"success": false,
			"message": "Internal Server Error",
			"data":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"success": true,
		"message": "Data successfully retrieved",
		"data":    provinces,
	})
}
