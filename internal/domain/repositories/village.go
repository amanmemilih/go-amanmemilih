package repositories

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type VillageRepository interface {
	FindAll(ctx context.Context, subdistrictId int) ([]*entities.Village, error)
}
