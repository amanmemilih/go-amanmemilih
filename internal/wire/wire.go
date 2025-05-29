//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/infrastructure/blockchain/icp"
	"github.com/zinct/amanmemilih/internal/infrastructure/clients/wordie"
	"github.com/zinct/amanmemilih/internal/infrastructure/ipfs/pinata"
	candidatRepo "github.com/zinct/amanmemilih/internal/infrastructure/repositories/candidat"
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

func InitializeDocumentController(cfg *config.Config, log *logger.Logger) (*controllers.DocumentController, error) {
	wire.Build(
		icp.NewClient,
		pinata.NewPinata,
		usecases.NewDocumentUsecase,
		controllers.NewDocumentController,
	)
	return nil, nil
}

func InitializePresidentialCandidatController(db *sql.DB, cfg *config.Config, log *logger.Logger) *controllers.PresidentialCandidatController {
	wire.Build(
		candidatRepo.NewPresidentialCandidatRepositoryMysql,
		usecases.NewPresidentialCandidatUsecase,
		controllers.NewPresidentialCandidatController,
	)

	return nil
}

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
