package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
	"github.com/zinct/amanmemilih/internal/interface/delivery/http/v1/response"
	"github.com/zinct/amanmemilih/pkg/logger"
)

type PresidentialCandidatController struct {
	uc  usecases.PresidentialCandidatUsecase
	cfg *config.Config
	log *logger.Logger
}

func NewPresidentialCandidatController(uc usecases.PresidentialCandidatUsecase, cfg *config.Config, log *logger.Logger) *PresidentialCandidatController {
	return &PresidentialCandidatController{
		uc:  uc,
		cfg: cfg,
		log: log,
	}
}

func (c *PresidentialCandidatController) FindAll(ctx *gin.Context) {
	candidates, err := c.uc.FindAll(ctx)
	if err != nil {
		response.JSONError(ctx, c.cfg, c.log, err)
		return
	}

	response.JSONSuccess(ctx, "Data successfully retrieved", candidates)
}

func (c *PresidentialCandidatController) Summary(ctx *gin.Context) {
	panic("not implemented")
}
