package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
	"github.com/zinct/amanmemilih/pkg/logger"
)

type VillageController struct {
	uc  usecases.VillageUsecase
	cfg *config.Config
	log *logger.Logger
}

func NewVillageController(villageUsecase usecases.VillageUsecase, cfg *config.Config, log *logger.Logger) *VillageController {
	return &VillageController{uc: villageUsecase, cfg: cfg, log: log}
}

func (c *VillageController) FindAll(ctx *gin.Context) {
	subdistrictId, err := strconv.Atoi(ctx.Param("subdistrictId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"success": false,
			"message": "Bad Request",
		})
		return
	}

	villages, err := c.uc.FindAll(ctx, subdistrictId)
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
		"data":    villages,
	})
}
