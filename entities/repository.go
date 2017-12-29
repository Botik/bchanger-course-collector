package entities

import (
	"github.com/jackc/pgx"
)

type Repository struct {
	conn   *pgx.Conn
	entity Entity
}

func (r *Repository) GetAll() ([]Entity, error) {
	rows, err := r.conn.Query("SELECT * FROM \"" + r.entity.GetTableName() + "\"")

	if nil != err {
		return nil, err
	}

	result := make([]Entity, 0, 128)

	for rows.Next() {
		entity, err := r.entity.Decode(rows)

		if nil != err {
			return nil, err
		}

		result = append(result, entity)
	}

	return result[:], nil
}
