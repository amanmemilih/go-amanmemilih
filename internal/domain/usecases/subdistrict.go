package usecases

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type SubdistrictUsecase interface {
	FindAll(ctx context.Context, districtId int) ([]*entities.Subdistrict, error)
}
