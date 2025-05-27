package usecases

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/repositories"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
)

type PresidentialCandidatUsecase struct {
	repo repositories.PresidentialCandidatRepository
}

func NewPresidentialCandidatUsecase(repo repositories.PresidentialCandidatRepository) usecases.PresidentialCandidatUsecase {
	return &PresidentialCandidatUsecase{repo: repo}
}

func (u *PresidentialCandidatUsecase) FindAll(ctx context.Context) ([]*entities.PresidentialCandidate, error) {
	candidates, err := u.repo.FindAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "PresidentialCandidatUsecase.FindAll")
	}
	return candidates, nil
}

func (u *PresidentialCandidatUsecase) Summary(ctx context.Context) ([]*entities.PresidentialCandidatSummary, error) {
	panic("not implemented")
}
