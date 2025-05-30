package http

import (
	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/interface/delivery/http/v1/controllers"
	"github.com/zinct/amanmemilih/internal/interface/delivery/http/v1/middleware"
	"github.com/zinct/amanmemilih/pkg/jwt"
	"github.com/zinct/amanmemilih/pkg/logger"

	"github.com/gin-gonic/gin"
)

type RouterOption struct {
	AuthController        *controllers.AuthController
	ProvinceController    *controllers.ProvinceController
	DistrictController    *controllers.DistrictController
	SubdistrictController *controllers.SubdistrictController
	VillageController     *controllers.VillageController
	CandidatController    *controllers.PresidentialCandidatController
	DocumentController    *controllers.DocumentController
}

func RegisterMiddleware(router *gin.Engine, cfg *config.Config, log *logger.Logger) {
	router.Use(gin.Logger())
	router.Use(middleware.Recovery(cfg, log))
}

func RegisterRoutes(router *gin.Engine, opts RouterOption, cfg *config.Config, log *logger.Logger, jm *jwt.JWTManager) *gin.Engine {
	{
		// BPS
		bps := router.Group("/bps")
		bps.GET("/province", opts.ProvinceController.FindAll)
		bps.GET("/district/:provinceId", opts.DistrictController.FindAll)
		bps.GET("/subdistrict/:districtId", opts.SubdistrictController.FindAll)
		bps.GET("/village/:subdistrictId", opts.VillageController.FindAll)
		bps.GET("/tps/:villageId", opts.DocumentController.GetDocumentUser)

		// Auth
		router.POST("/login", opts.AuthController.Login)
		router.POST("/register", opts.AuthController.Register)
		router.POST("/recovery-key", opts.AuthController.GeneratePhrase)
		router.POST("/forgot-password", opts.AuthController.ChangePassword)
		router.GET("/check-credentials", middleware.JWTAuthMiddleware(jm, cfg, log), opts.AuthController.CheckCredential)
		router.POST("/logout", middleware.JWTAuthMiddleware(jm, cfg, log), opts.AuthController.Logout)

		// Candidat
		router.GET("/presidential-candidats", middleware.JWTAuthMiddleware(jm, cfg, log), opts.CandidatController.FindAll)
		router.GET("/presidential-candidats/summary", middleware.JWTAuthMiddleware(jm, cfg, log), opts.DocumentController.Summary)

		router.GET("/documents", middleware.JWTAuthMiddleware(jm, cfg, log), opts.DocumentController.FindAll)
		router.POST("/documents", middleware.JWTAuthMiddleware(jm, cfg, log), opts.DocumentController.Create)
		router.GET("/documents/:id", middleware.JWTAuthMiddleware(jm, cfg, log), opts.DocumentController.Find)
		router.POST("/documents/:id/verified", middleware.JWTAuthMiddleware(jm, cfg, log), opts.DocumentController.Verify)

		router.GET("/dashboard", middleware.JWTAuthMiddleware(jm, cfg, log), opts.DocumentController.Dashboard)
	}

	return router
}
