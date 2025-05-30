package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/repositories"
	apperr "github.com/zinct/amanmemilih/internal/errors"
)

type UserRepositoryMysql struct {
	db *sql.DB
}

func NewUserRepositoryMysql(db *sql.DB) repositories.UserRepository {
	return &UserRepositoryMysql{db: db}
}

func (r *UserRepositoryMysql) UpdatePasswordByID(ctx context.Context, userID int, password string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET password = ? WHERE id = ?", password, userID)
	if err != nil {
		return fmt.Errorf("internal/infrastructure/repositories/user/mysql - UpdatePasswordByID - ExecContext: %w", err)
	}
	return nil
}

func (r *UserRepositoryMysql) UpdateUsernameVerifiedAtByID(ctx context.Context, userID int) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET username_verified_at = NOW() WHERE id = ?", userID)
	if err != nil {
		return fmt.Errorf("internal/infrastructure/repositories/user/mysql - UpdateUsernameVerifiedAtByID - ExecContext: %w", err)
	}
	return nil
}

func (r *UserRepositoryMysql) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	row := r.db.QueryRowContext(ctx, `
	SELECT 
		users.id, 
		username, 
		username_verified_at, 
		password, 
		village_id, 
		address, 
		users.created_at, 
		users.updated_at, 
		villages.name as village,
		provinces.name as province, 
		districts.name as district, 
		subdistricts.name as subdistrict, 
		CONCAT(address, ', ', villages.name, ', ', districts.name, ', ', subdistricts.name, ', ', provinces.name) as region 
	FROM users 
	JOIN villages ON users.village_id = villages.id 
	JOIN subdistricts ON villages.subdistrict_id = subdistricts.id 
	JOIN districts ON subdistricts.district_id = districts.id 
	JOIN provinces ON districts.province_id = provinces.id 
	WHERE username = ?`, username)
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("internal/infrastructure/repositories/user/mysql - FindByUsername - QueryRowContext: %w", err)
	}

	var user entities.User
	if err := row.Scan(&user.Id, &user.Username, &user.UsernameVerifiedAt, &user.Password, &user.VillageId, &user.Address, &user.CreatedAt, &user.UpdatedAt, &user.Village, &user.Province, &user.District, &user.Subdistrict, &user.Region); err != nil {
		return nil, fmt.Errorf("internal/infrastructure/repositories/user/mysql - FindByID - Scan: %w", err)
	}

	return &user, nil
}

func (r *UserRepositoryMysql) FindByID(ctx context.Context, id int) (*entities.User, error) {
	row := r.db.QueryRowContext(ctx, `
	SELECT 
		users.id, 
		username, 
		username_verified_at, 
		password, 
		village_id, 
		address, 
		users.created_at, 
		users.updated_at, 
		villages.name as village,
		provinces.name as province, 
		districts.name as district, 
		subdistricts.name as subdistrict, 
		CONCAT(address, ', ', villages.name, ', ', districts.name, ', ', subdistricts.name, ', ', provinces.name) as region 
	FROM users 
	JOIN villages ON users.village_id = villages.id 
	JOIN subdistricts ON villages.subdistrict_id = subdistricts.id 
	JOIN districts ON subdistricts.district_id = districts.id 
	JOIN provinces ON districts.province_id = provinces.id 
	WHERE users.id = ?`, id)
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("internal/infrastructure/repositories/user/mysql - FindByID - QueryRowContext: %w", err)
	}

	var user entities.User
	if err := row.Scan(&user.Id, &user.Username, &user.UsernameVerifiedAt, &user.Password, &user.VillageId, &user.Address, &user.CreatedAt, &user.UpdatedAt, &user.Village, &user.Province, &user.District, &user.Subdistrict, &user.Region); err != nil {
		return nil, fmt.Errorf("internal/infrastructure/repositories/user/mysql - FindByID - Scan: %w", err)
	}

	return &user, nil
}

func (r *UserRepositoryMysql) CreatePhrase(ctx context.Context, username string, phrase *entities.Phrase) error {
	// Check if username exists
	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM password_reset_tokens WHERE username = ?)", username).Scan(&exists)
	if err != nil {
		return fmt.Errorf("internal/infrastructure/repositories/user/mysql - CreatePhrase - QueryRowContext: %w", err)
	}

	if exists {
		// Update existing record
		_, err = r.db.ExecContext(ctx, `
			UPDATE password_reset_tokens SET
				phrase_1 = ?, phrase_2 = ?, phrase_3 = ?, phrase_4 = ?,
				phrase_5 = ?, phrase_6 = ?, phrase_7 = ?, phrase_8 = ?,
				phrase_9 = ?, phrase_10 = ?, phrase_11 = ?, phrase_12 = ?
			WHERE username = ?`,
			phrase.Phrase1, phrase.Phrase2, phrase.Phrase3, phrase.Phrase4,
			phrase.Phrase5, phrase.Phrase6, phrase.Phrase7, phrase.Phrase8,
			phrase.Phrase9, phrase.Phrase10, phrase.Phrase11, phrase.Phrase12,
			username)
	} else {
		// Insert new record
		_, err = r.db.ExecContext(ctx, `
			INSERT INTO password_reset_tokens (
				username, phrase_1, phrase_2, phrase_3, phrase_4,
				phrase_5, phrase_6, phrase_7, phrase_8, phrase_9,
				phrase_10, phrase_11, phrase_12
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			username, phrase.Phrase1, phrase.Phrase2, phrase.Phrase3, phrase.Phrase4,
			phrase.Phrase5, phrase.Phrase6, phrase.Phrase7, phrase.Phrase8, phrase.Phrase9,
			phrase.Phrase10, phrase.Phrase11, phrase.Phrase12)
	}

	if err != nil {
		return fmt.Errorf("internal/infrastructure/repositories/user/mysql - CreatePhrase - ExecContext: %w", err)
	}
	return nil
}

func (r *UserRepositoryMysql) FindByPhrase(ctx context.Context, phrase1, phrase2, phrase3, phrase4, phrase5, phrase6, phrase7, phrase8, phrase9, phrase10, phrase11, phrase12 string) (*entities.Phrase, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, username, phrase_1, phrase_2, phrase_3, phrase_4, phrase_5, phrase_6, phrase_7, phrase_8, phrase_9, phrase_10, phrase_11, phrase_12 FROM password_reset_tokens WHERE phrase_1 = ? AND phrase_2 = ? AND phrase_3 = ? AND phrase_4 = ? AND phrase_5 = ? AND phrase_6 = ? AND phrase_7 = ? AND phrase_8 = ? AND phrase_9 = ? AND phrase_10 = ? AND phrase_11 = ? AND phrase_12 = ?", phrase1, phrase2, phrase3, phrase4, phrase5, phrase6, phrase7, phrase8, phrase9, phrase10, phrase11, phrase12)
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("internal/infrastructure/repositories/user/mysql - FindByPhrase - QueryRowContext: %w", err)
	}

	var phrase entities.Phrase
	if err := row.Scan(&phrase.Id, &phrase.Username, &phrase.Phrase1, &phrase.Phrase2, &phrase.Phrase3, &phrase.Phrase4, &phrase.Phrase5, &phrase.Phrase6, &phrase.Phrase7, &phrase.Phrase8, &phrase.Phrase9, &phrase.Phrase10, &phrase.Phrase11, &phrase.Phrase12); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperr.NewAPPError(422, "Phrase is invalid", apperr.AppError, nil)
		}
		return nil, fmt.Errorf("internal/infrastructure/repositories/user/mysql - FindByPhrase - Scan: %w", err)
	}

	return &phrase, nil
}

func (r *UserRepositoryMysql) FindByVillageID(ctx context.Context, villageID int) ([]entities.User, error) {
	rows, err := r.db.QueryContext(ctx, `
	SELECT 
		users.id, 
		username, 
		username_verified_at, 
		password, 
		village_id, 
		address, 
		users.created_at, 
		users.updated_at, 
		villages.name as village,
		provinces.name as province, 
		districts.name as district, 
		subdistricts.name as subdistrict, 
		CONCAT(address, ', ', villages.name, ', ', districts.name, ', ', subdistricts.name, ', ', provinces.name) as region 
	FROM users 
	JOIN villages ON users.village_id = villages.id 
	JOIN subdistricts ON villages.subdistrict_id = subdistricts.id 
	JOIN districts ON subdistricts.district_id = districts.id 
	JOIN provinces ON districts.province_id = provinces.id 
	WHERE users.village_id = ?`, villageID)
	if err != nil {
		return nil, fmt.Errorf("internal/infrastructure/repositories/user/mysql - FindByVillageID - QueryContext: %w", err)
	}

	var users []entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.Id, &user.Username, &user.UsernameVerifiedAt, &user.Password, &user.VillageId, &user.Address, &user.CreatedAt, &user.UpdatedAt, &user.Village, &user.Province, &user.District, &user.Subdistrict, &user.Region)
		if err != nil {
			return nil, fmt.Errorf("internal/infrastructure/repositories/user/mysql - FindByVillageID - Scan: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
