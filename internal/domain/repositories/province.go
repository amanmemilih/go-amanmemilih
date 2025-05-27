package repositories

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type ProvinceRepository interface {
	FindAll(ctx context.Context) ([]*entities.Province, error)
}
