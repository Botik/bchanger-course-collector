package entities

import "github.com/jackc/pgx"

type Course struct {
	Entity
	From uint8
	To   uint8
	Rate float64
}

func (c Course) Decode(r *pgx.Rows) (Entity, error) {
	return c, r.Scan(&c.From, &c.To, &c.Rate)
}

func (c *Course) getTableName() string {
	return "Course"
}
