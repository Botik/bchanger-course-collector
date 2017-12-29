package entities

import "github.com/jackc/pgx"

type Currency struct {
	Entity
	ID       uint8
	Name     string
	Charcode string
	Symbol   string
}

func (c Currency) Decode(r *pgx.Rows) (Entity, error) {
	var strings [3][]byte

	err := r.Scan(&c.ID, &strings[0], &strings[1], &strings[2])

	if nil != err {
		return c, err
	}

	c.Name = decodeEnmptyString(strings[0])
	c.Charcode = decodeEnmptyString(strings[1])
	c.Symbol = decodeEnmptyString(strings[2])

	return c, err
}

func (c Currency) GetTableName() string {
	return "Currency"
}
