package usecases

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type DocumentUsecase interface {
	FindAll(ctx context.Context) ([]*entities.Document, error)
	Find(ctx context.Context, id int, electionType string) (*entities.Document, error)
	Verify(ctx context.Context, id int, electionType string) error
	Create(ctx context.Context, userId int, electionType string, votes []entities.DocumentVote, documents []string, documentNames []string) error
	Summary(ctx context.Context) (interface{}, error)
}
