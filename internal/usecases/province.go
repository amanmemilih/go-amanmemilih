package usecases

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/repositories"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
)

type ProvinceUsecase struct {
	repo repositories.ProvinceRepository
}

func NewProvinceUsecase(repo repositories.ProvinceRepository) usecases.ProvinceUsecase {
	return &ProvinceUsecase{repo: repo}
}

func (uc *ProvinceUsecase) FindAll(ctx context.Context) ([]*entities.Province, error) {
	return uc.repo.FindAll(ctx)
}
