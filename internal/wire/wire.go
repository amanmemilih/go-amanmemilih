//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/infrastructure/clients/wordie"
	districtRepo "github.com/zinct/amanmemilih/internal/infrastructure/repositories/district"
	provinceRepo "github.com/zinct/amanmemilih/internal/infrastructure/repositories/province"
	subdistrictRepo "github.com/zinct/amanmemilih/internal/infrastructure/repositories/subdistrict"
	userRepo "github.com/zinct/amanmemilih/internal/infrastructure/repositories/user"
	villageRepo "github.com/zinct/amanmemilih/internal/infrastructure/repositories/village"
	"github.com/zinct/amanmemilih/internal/interface/delivery/http/v1/controllers"
	"github.com/zinct/amanmemilih/internal/usecases"
	"github.com/zinct/amanmemilih/pkg/jwt"
	"github.com/zinct/amanmemilih/pkg/logger"
)

func InitializeAuthController(db *sql.DB, cfg *config.Config, log *logger.Logger, jwtManager *jwt.JWTManager) *controllers.AuthController {
	wire.Build(
		wordie.NewClient,
		userRepo.NewUserRepositoryMysql,
		usecases.NewAuthUsecase,
		controllers.NewAuthController,
	)
	return nil
}

func InitializeProvinceController(db *sql.DB, cfg *config.Config, log *logger.Logger) *controllers.ProvinceController {
	wire.Build(
		provinceRepo.NewProvinceRepositoryMysql,
		usecases.NewProvinceUsecase,
		controllers.NewProvinceController,
	)

	return nil
}

func InitializeDistrictController(db *sql.DB, cfg *config.Config, log *logger.Logger) *controllers.DistrictController {
	wire.Build(
		districtRepo.NewDistrictRepositoryMysql,
		usecases.NewDistrictUsecase,
		controllers.NewDistrictController,
	)

	return nil
}

func InitializeSubdistrictController(db *sql.DB, cfg *config.Config, log *logger.Logger) *controllers.SubdistrictController {
	wire.Build(
		subdistrictRepo.NewSubdistrictRepositoryMysql,
		usecases.NewSubdistrictUsecase,
		controllers.NewSubdistrictController,
	)

	return nil
}

func InitializeVillageController(db *sql.DB, cfg *config.Config, log *logger.Logger) *controllers.VillageController {
	wire.Build(
		villageRepo.NewVillageRepositoryMysql,
		usecases.NewVillageUsecase,
		controllers.NewVillageController,
	)

	return nil
}
