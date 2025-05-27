package usecases

import (
	"context"

	"github.com/zinct/amanmemilih/internal/domain/entities"
)

type AuthUsecase interface {
	Login(ctx context.Context, username, password string) (*entities.User, string, error)
	Register(ctx context.Context, username, password, phrase1, phrase2, phrase3, phrase4, phrase5, phrase6, phrase7, phrase8, phrase9, phrase10, phrase11, phrase12 string) error
	GeneratePhrase(ctx context.Context, username string) (*entities.Phrase, error)
	CheckCredential(ctx context.Context, userID int) (*entities.User, error)
	Logout(ctx context.Context) error
	ChangePassword(ctx context.Context, password, phrase1, phrase2, phrase3, phrase4, phrase5, phrase6, phrase7, phrase8, phrase9, phrase10, phrase11, phrase12 string) error
}
