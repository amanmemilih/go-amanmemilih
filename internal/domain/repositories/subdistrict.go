package repositories

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type SubdistrictRepository interface {
	FindAll(ctx context.Context, districtId int) ([]*entities.Subdistrict, error)
}
