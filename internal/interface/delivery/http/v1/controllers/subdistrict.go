package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
	"github.com/zinct/amanmemilih/pkg/logger"
)

type SubdistrictController struct {
	uc  usecases.SubdistrictUsecase
	cfg *config.Config
	log *logger.Logger
}

func NewSubdistrictController(subdistrictUsecase usecases.SubdistrictUsecase, cfg *config.Config, log *logger.Logger) *SubdistrictController {
	return &SubdistrictController{uc: subdistrictUsecase, cfg: cfg, log: log}
}

func (c *SubdistrictController) FindAll(ctx *gin.Context) {
	districtId, err := strconv.Atoi(ctx.Param("districtId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"success": false,
			"message": "Bad Request",
		})
		return
	}

	subdistricts, err := c.uc.FindAll(ctx, districtId)
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
		"data":    subdistricts,
	})
}
