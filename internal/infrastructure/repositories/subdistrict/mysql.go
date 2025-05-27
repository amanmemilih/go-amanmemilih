package province

import (
	"context"
	"database/sql"

	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/repositories"
)

type SubdistrictRepositoryMysql struct {
	db *sql.DB
}

func NewSubdistrictRepositoryMysql(db *sql.DB) repositories.SubdistrictRepository {
	return &SubdistrictRepositoryMysql{db: db}
}

func (r *SubdistrictRepositoryMysql) FindAll(ctx context.Context, districtId int) ([]*entities.Subdistrict, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM subdistricts WHERE district_id = ?", districtId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subdistricts := []*entities.Subdistrict{}
	for rows.Next() {
		var subdistrict entities.Subdistrict
		err = rows.Scan(&subdistrict.Id, &subdistrict.Code, &subdistrict.Name, &subdistrict.DistrictId, &subdistrict.CreatedAt, &subdistrict.UpdatedAt)
		if err != nil {
			return nil, err
		}
		subdistricts = append(subdistricts, &subdistrict)
	}

	return subdistricts, nil
}
