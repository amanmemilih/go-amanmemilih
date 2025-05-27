package usecases

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type DistrictUsecase interface {
	FindAll(ctx context.Context, provinceId int) ([]*entities.District, error)
}
