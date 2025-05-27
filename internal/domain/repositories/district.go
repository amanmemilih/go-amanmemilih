package repositories

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type DistrictRepository interface {
	FindAll(ctx context.Context, provinceId int) ([]*entities.District, error)
}
