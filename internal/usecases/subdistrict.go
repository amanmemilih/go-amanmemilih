package usecases

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/repositories"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
)

type SubdistrictUsecase struct {
	repo repositories.SubdistrictRepository
}

func NewSubdistrictUsecase(repo repositories.SubdistrictRepository) usecases.SubdistrictUsecase {
	return &SubdistrictUsecase{repo: repo}
}

func (uc *SubdistrictUsecase) FindAll(ctx context.Context, districtId int) ([]*entities.Subdistrict, error) {
	return uc.repo.FindAll(ctx, districtId)
}
