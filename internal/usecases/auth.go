package usecases

import (
	"context"
	"database/sql"

	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/interfaces"
	"github.com/zinct/amanmemilih/internal/domain/repositories"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
	apperr "github.com/zinct/amanmemilih/internal/errors"
	"github.com/zinct/amanmemilih/internal/utils"
	"github.com/zinct/amanmemilih/pkg/jwt"
)

type AuthUsecase struct {
	userRepo   repositories.UserRepository
	wordClient interfaces.WordClient
	jwtManager *jwt.JWTManager
}

func NewAuthUsecase(userRepo repositories.UserRepository, jwtManager *jwt.JWTManager, wordClient interfaces.WordClient) usecases.AuthUsecase {
	return &AuthUsecase{userRepo: userRepo, jwtManager: jwtManager, wordClient: wordClient}
}

func (u *AuthUsecase) Login(ctx context.Context, username, password string) (*entities.User, string, error) {
	user, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, "", apperr.NewUnauthorizedError("invalid credential", nil)
	}

	if user.UsernameVerifiedAt == nil {
		return nil, "", apperr.NewAPPError(422, "Your account not registered", apperr.AppError, nil)
	}

	if !utils.CheckPassword(password, *user.Password) {
		return nil, "", apperr.NewUnauthorizedError("invalid credential", nil)
	}

	token, err := u.jwtManager.GenerateJWT(user.Id)
	if err != nil {
		return nil, "", apperr.NewInternalError("failed to generate token", nil)
	}

	return user, token, nil
}

func (u *AuthUsecase) Register(ctx context.Context, username, password, phrase1, phrase2, phrase3, phrase4, phrase5, phrase6, phrase7, phrase8, phrase9, phrase10, phrase11, phrase12 string) error {
	panic("not implemented")
}

func (u *AuthUsecase) GeneratePhrase(ctx context.Context, username string) (*entities.Phrase, error) {
	user, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, apperr.NewAPPError(422, "User not found", apperr.AppError, nil)
	}

	if user.UsernameVerifiedAt != nil {
		return nil, apperr.NewAPPError(422, "Your account already registered", apperr.AppError, nil)
	}

	words, err := u.wordClient.GetRandomWords(12)
	if err != nil {
		return nil, err
	}

	phrase := entities.Phrase{
		Id:       user.Id,
		Username: username,
		Phrase1:  words[0],
		Phrase2:  words[1],
		Phrase3:  words[2],
		Phrase4:  words[3],
		Phrase5:  words[4],
		Phrase6:  words[5],
		Phrase7:  words[6],
		Phrase8:  words[7],
		Phrase9:  words[8],
		Phrase10: words[9],
		Phrase11: words[10],
		Phrase12: words[11],
	}

	err = u.userRepo.CreatePhrase(ctx, user.Username, &phrase)
	if err != nil {
		return nil, err
	}

	return &phrase, nil
}

func (u *AuthUsecase) CheckCredential(ctx context.Context, userID int) (*entities.User, error) {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, apperr.NewUnauthorizedError("invalid credential", nil)
	}

	return user, nil
}

func (u *AuthUsecase) Logout(ctx context.Context) error {
	panic("not implemented")
}

func (u *AuthUsecase) ChangePassword(ctx context.Context, password, phrase1, phrase2, phrase3, phrase4, phrase5, phrase6, phrase7, phrase8, phrase9, phrase10, phrase11, phrase12 string) error {
	type phraseResult struct {
		phrase *entities.Phrase
		err    error
	}

	type hashResult struct {
		hash string
		err  error
	}

	phraseChan := make(chan phraseResult, 1)
	hashChan := make(chan hashResult, 1)

	go func() {
		phrase, err := u.userRepo.FindByPhrase(ctx, phrase1, phrase2, phrase3, phrase4, phrase5, phrase6, phrase7, phrase8, phrase9, phrase10, phrase11, phrase12)
		phraseChan <- phraseResult{phrase: phrase, err: err}
	}()

	go func() {
		hash, err := utils.HashPassword(password)
		hashChan <- hashResult{hash: hash, err: err}
	}()

	var phrase phraseResult
	var hash hashResult

	select {
	case phrase = <-phraseChan:
	case <-ctx.Done():
		return ctx.Err()
	}

	select {
	case hash = <-hashChan:
	case <-ctx.Done():
		return ctx.Err()
	}

	if phrase.err != nil {
		if phrase.err == sql.ErrNoRows {
			return apperr.NewAPPError(422, "Phrase not invalid", apperr.AppError, nil)
		}
		return phrase.err
	}

	if hash.err != nil {
		return hash.err
	}

	err := u.userRepo.ChangePassword(ctx, phrase.phrase.Id, hash.hash)
	if err != nil {
		return err
	}

	return nil
}
