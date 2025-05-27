package province

import (
	"context"
	"database/sql"

	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/repositories"
)

type ProvinceRepositoryMysql struct {
	db *sql.DB
}

func NewProvinceRepositoryMysql(db *sql.DB) repositories.ProvinceRepository {
	return &ProvinceRepositoryMysql{db: db}
}

func (r *ProvinceRepositoryMysql) FindAll(ctx context.Context) ([]*entities.Province, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM provinces")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	provinces := []*entities.Province{}
	for rows.Next() {
		var province entities.Province
		err = rows.Scan(&province.Id, &province.Code, &province.Name, &province.CreatedAt, &province.UpdatedAt)
		if err != nil {
			return nil, err
		}
		provinces = append(provinces, &province)
	}

	return provinces, nil
}
