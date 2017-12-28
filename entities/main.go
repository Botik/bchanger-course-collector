package entities

import "github.com/jackc/pgx"

var (
	CurrencyType = iota
)

type DBDecoder interface {
	Decode(r *pgx.Rows) error
}

func Unmarshal(r *pgx.Rows, d DBDecoder) error {
	return d.Decode(r)
}

func decodeEnmptyString(b []byte) string {
	if nil == b {
		return ""
	}

	return string(b)
}
