package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
	apperr "github.com/zinct/amanmemilih/internal/errors"
	"github.com/zinct/amanmemilih/internal/interface/delivery/http/v1/presenter"
	"github.com/zinct/amanmemilih/internal/interface/delivery/http/v1/request"
	"github.com/zinct/amanmemilih/internal/interface/delivery/http/v1/response"
	"github.com/zinct/amanmemilih/internal/utils"
	"github.com/zinct/amanmemilih/pkg/logger"
)

type AuthController struct {
	authUsecase usecases.AuthUsecase
	cfg         *config.Config
	log         *logger.Logger
}

func NewAuthController(authUsecase usecases.AuthUsecase, cfg *config.Config, log *logger.Logger) *AuthController {
	return &AuthController{
		authUsecase: authUsecase,
		cfg:         cfg,
		log:         log,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest
	if err := ctx.ShouldBind(&loginRequest); err != nil {
		err := apperr.NewValidationError("Invalid request format", utils.FormatValidationError(err))
		response.JSONError(ctx, c.cfg, c.log, err)
		return
	}

	user, token, err := c.authUsecase.Login(ctx, loginRequest.Username, loginRequest.Password)
	if err != nil {
		response.JSONError(ctx, c.cfg, c.log, err)
		return
	}

	response.JSONSuccess(ctx, "Login successful", presenter.LoginResponse{
		User: presenter.UserLoginResponse{
			Id:          user.Id,
			Username:    user.Username,
			Address:     user.Address,
			Village:     user.Village,
			Province:    user.Province,
			District:    user.District,
			Subdistrict: user.Subdistrict,
			Region:      user.Region,
			CreatedAt:   user.CreatedAt,
		},
		Token: token,
	})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var request request.RegisterRequest
	if err := ctx.ShouldBind(&request); err != nil {
		err := apperr.NewValidationError("Invalid request format", utils.FormatValidationError(err))
		response.JSONError(ctx, c.cfg, c.log, err)
		return
	}

	err := c.authUsecase.Register(ctx, request.Username, request.Password, request.Phrase1, request.Phrase2, request.Phrase3, request.Phrase4, request.Phrase5, request.Phrase6, request.Phrase7, request.Phrase8, request.Phrase9, request.Phrase10, request.Phrase11, request.Phrase12)
	if err != nil {
		response.JSONError(ctx, c.cfg, c.log, err)
		return
	}

	response.JSONSuccess(ctx, "Register successful", nil)
}

func (c *AuthController) GeneratePhrase(ctx *gin.Context) {
	var request request.RecoveryKeyRequest
	if err := ctx.ShouldBind(&request); err != nil {
		err := apperr.NewValidationError("Invalid request format", utils.FormatValidationError(err))
		response.JSONError(ctx, c.cfg, c.log, err)
		return
	}

	phrase, err := c.authUsecase.GeneratePhrase(ctx, request.Username)
	if err != nil {
		response.JSONError(ctx, c.cfg, c.log, err)
		return
	}

	response.JSONSuccess(ctx, "Generate phrase successful", phrase)
}

func (c *AuthController) CheckCredential(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	token := ctx.GetString("token")

	user, err := c.authUsecase.CheckCredential(ctx, userID)
	if err != nil {
		response.JSONError(ctx, c.cfg, c.log, err)
		return
	}

	response.JSONSuccess(ctx, "Check credential successful", presenter.LoginResponse{
		User: presenter.UserLoginResponse{
			Id:          user.Id,
			Username:    user.Username,
			Address:     user.Address,
			Village:     user.Village,
			Province:    user.Province,
			District:    user.District,
			Subdistrict: user.Subdistrict,
			Region:      user.Region,
			CreatedAt:   user.CreatedAt,
		},
		Token: token,
	})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	panic("not implemented")
}

func (c *AuthController) ChangePassword(ctx *gin.Context) {
	var request request.ChangePasswordRequest
	if err := ctx.ShouldBind(&request); err != nil {
		err := apperr.NewValidationError("Invalid request format", utils.FormatValidationError(err))
		response.JSONError(ctx, c.cfg, c.log, err)
		return
	}

	err := c.authUsecase.ChangePassword(ctx, request.Password, request.Phrase1, request.Phrase2, request.Phrase3, request.Phrase4, request.Phrase5, request.Phrase6, request.Phrase7, request.Phrase8, request.Phrase9, request.Phrase10, request.Phrase11, request.Phrase12)
	if err != nil {
		response.JSONError(ctx, c.cfg, c.log, err)
		return
	}

	response.JSONSuccess(ctx, "Change password successful", nil)
}
