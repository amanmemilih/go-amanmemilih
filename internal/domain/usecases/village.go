package usecases

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type VillageUsecase interface {
	FindAll(ctx context.Context, subdistrictId int) ([]*entities.Village, error)
}
