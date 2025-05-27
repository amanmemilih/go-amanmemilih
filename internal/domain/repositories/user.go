package repositories

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type UserRepository interface {
	UpdatePasswordByID(ctx context.Context, userID int, password string) error
	UpdateUsernameVerifiedAtByID(ctx context.Context, userID int) error
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
	FindByID(ctx context.Context, id int) (*entities.User, error)
	CreatePhrase(ctx context.Context, username string, phrase *entities.Phrase) error
	FindByPhrase(ctx context.Context, phrase1, phrase2, phrase3, phrase4, phrase5, phrase6, phrase7, phrase8, phrase9, phrase10, phrase11, phrase12 string) (*entities.Phrase, error)
}
