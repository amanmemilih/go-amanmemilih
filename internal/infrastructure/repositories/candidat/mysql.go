package candidat

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/repositories"
)

type PresidentialCandidatRepositoryMysql struct {
	db *sql.DB
}

func NewPresidentialCandidatRepositoryMysql(db *sql.DB) repositories.PresidentialCandidatRepository {
	return &PresidentialCandidatRepositoryMysql{db: db}
}

func (r *PresidentialCandidatRepositoryMysql) FindAll(ctx context.Context) ([]*entities.PresidentialCandidate, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, no, name, image FROM presidential_candidats")
	if err != nil {
		return nil, errors.Wrap(err, "PresidentialCandidatRepositoryMysql.FindAll.QueryContext")
	}
	defer rows.Close()

	var candidates []*entities.PresidentialCandidate
	for rows.Next() {
		var candidate entities.PresidentialCandidate
		if err := rows.Scan(&candidate.Id, &candidate.No, &candidate.Name, &candidate.Image); err != nil {
			return nil, errors.Wrap(err, "PresidentialCandidatRepositoryMysql.FindAll.Scan")
		}
		candidates = append(candidates, &candidate)
	}
	return candidates, nil
}
