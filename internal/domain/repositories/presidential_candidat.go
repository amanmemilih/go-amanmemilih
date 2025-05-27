package repositories

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type PresidentialCandidatRepository interface {
	FindAll(ctx context.Context) ([]*entities.PresidentialCandidate, error)
}
