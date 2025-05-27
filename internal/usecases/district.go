package usecases

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/repositories"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
)

type DistrictUsecase struct {
	repo repositories.DistrictRepository
}

func NewDistrictUsecase(repo repositories.DistrictRepository) usecases.DistrictUsecase {
	return &DistrictUsecase{repo: repo}
}

func (uc *DistrictUsecase) FindAll(ctx context.Context, provinceId int) ([]*entities.District, error) {
	return uc.repo.FindAll(ctx, provinceId)
}
