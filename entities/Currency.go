package entities

import "github.com/jackc/pgx"

type Currency struct {
	DBDecoder
	ID       uint8
	Name     string
	Charcode string
	Symbol   string
}

var tableName = "Currency"

func (c *Currency) Decode(r *pgx.Rows) error {
	var strings [3][]byte

	err := r.Scan(&c.ID, &strings[0], &strings[1], &strings[2])

	if nil != err {
		return err
	}

	c.Name = decodeEnmptyString(strings[0])
	c.Charcode = decodeEnmptyString(strings[1])
	c.Symbol = decodeEnmptyString(strings[2])

	return nil
}
