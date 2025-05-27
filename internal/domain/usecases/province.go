package usecases

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type ProvinceUsecase interface {
	FindAll(ctx context.Context) ([]*entities.Province, error)
}
