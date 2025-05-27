package usecases

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/repositories"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
)

type VillageUsecase struct {
	repo repositories.VillageRepository
}

func NewVillageUsecase(repo repositories.VillageRepository) usecases.VillageUsecase {
	return &VillageUsecase{repo: repo}
}

func (uc *VillageUsecase) FindAll(ctx context.Context, subdistrictId int) ([]*entities.Village, error) {
	return uc.repo.FindAll(ctx, subdistrictId)
}
