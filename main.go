package main

import (
	"fmt"
	"local/courseCollector/entities"

	"github.com/jackc/pgx"
)

func main() {
	connStr, _ := pgx.ParseURI("postgres://change:change@localhost/change")
	db, err := pgx.Connect(connStr)
	defer db.Close()

	if nil != err {
		panic(err)
	}

	storage := entities.Storage(db)
	currencyRepository := storage.GetRepository(entities.CurrencyType)
	rows, err := currencyRepository.GetAll()

	if nil != err {
		panic(err)
	}

	fmt.Printf("%+v", rows)
}
