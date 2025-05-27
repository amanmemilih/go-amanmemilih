package province

import (
	"context"
	"database/sql"

	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/repositories"
)

type VillageRepositoryMysql struct {
	db *sql.DB
}

func NewVillageRepositoryMysql(db *sql.DB) repositories.VillageRepository {
	return &VillageRepositoryMysql{db: db}
}

func (r *VillageRepositoryMysql) FindAll(ctx context.Context, subdistrictId int) ([]*entities.Village, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM villages WHERE subdistrict_id = ?", subdistrictId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	villages := []*entities.Village{}
	for rows.Next() {
		var village entities.Village
		err = rows.Scan(&village.Id, &village.Code, &village.Name, &village.SubdistrictId, &village.CreatedAt, &village.UpdatedAt)
		if err != nil {
			return nil, err
		}
		villages = append(villages, &village)
	}

	return villages, nil
}
