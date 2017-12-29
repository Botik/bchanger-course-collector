package entities

import (
	"errors"

	"github.com/jackc/pgx"
)

const (
	CourseType uint8 = iota
	CurrencyType
)

type Stor struct {
	conn *pgx.Conn
}

type Entity interface {
	GetTableName() string
	Decode(r *pgx.Rows) (Entity, error)
}

func Storage(conn *pgx.Conn) Stor {
	return Stor{conn}
}

func (s *Stor) GetRepository(repType uint8) (Repository, error) {
	switch repType {
	case CourseType:
		return Repository{s.conn, Course{}}, nil
	case CurrencyType:
		return Repository{s.conn, Currency{}}, nil
	}

	return Repository{}, errors.New("Undefined Repository")
}

func decodeEnmptyString(b []byte) string {
	if nil == b {
		return ""
	}

	return string(b)
}
