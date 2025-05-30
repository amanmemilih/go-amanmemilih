package usecases

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/interfaces"
)

type DocumentUsecase interface {
	FindAll(ctx context.Context, userId int) ([]interfaces.CheckDocumentResponse, error)
	Find(ctx context.Context, id int, electionType string) (*interfaces.PresidentialDocumentDetailResponse, error)
	Verify(ctx context.Context, id int, electionType string) error
	Create(ctx context.Context, userId int, electionType string, votes []entities.DocumentVote, documents []string, documentNames []string) error
	Summary(ctx context.Context) ([]interfaces.VotePercentage, error)
	Dashboard(ctx context.Context, userId int) (*interfaces.DashboardResponse, error)
}
