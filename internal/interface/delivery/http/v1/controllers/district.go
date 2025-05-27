package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
	"github.com/zinct/amanmemilih/pkg/logger"
)

type DistrictController struct {
	uc  usecases.DistrictUsecase
	cfg *config.Config
	log *logger.Logger
}

func NewDistrictController(districtUsecase usecases.DistrictUsecase, cfg *config.Config, log *logger.Logger) *DistrictController {
	return &DistrictController{uc: districtUsecase, cfg: cfg, log: log}
}

func (c *DistrictController) FindAll(ctx *gin.Context) {
	provinceId, err := strconv.Atoi(ctx.Param("provinceId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"success": false,
			"message": "Bad Request",
		})
		return
	}

	provinces, err := c.uc.FindAll(ctx, provinceId)
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
