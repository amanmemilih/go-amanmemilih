package province

import (
	"context"
	"database/sql"

	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/repositories"
)

type DistrictRepositoryMysql struct {
	db *sql.DB
}

func NewDistrictRepositoryMysql(db *sql.DB) repositories.DistrictRepository {
	return &DistrictRepositoryMysql{db: db}
}

func (r *DistrictRepositoryMysql) FindAll(ctx context.Context, provinceId int) ([]*entities.District, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM districts WHERE province_id = ?", provinceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	districts := []*entities.District{}
	for rows.Next() {
		var district entities.District
		err = rows.Scan(&district.Id, &district.Code, &district.Name, &district.ProvinceId, &district.CreatedAt, &district.UpdatedAt)
		if err != nil {
			return nil, err
		}
		districts = append(districts, &district)
	}

	return districts, nil
}
