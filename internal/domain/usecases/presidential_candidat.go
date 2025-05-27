package usecases

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type PresidentialCandidatUsecase interface {
	FindAll(ctx context.Context) ([]*entities.PresidentialCandidate, error)
	Summary(ctx context.Context) ([]*entities.PresidentialCandidatSummary, error)
}
