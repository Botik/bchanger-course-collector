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

	rows, err := db.Query("SELECT * FROM \"Currency\"")
	defer rows.Close()
	rows.
		println(len(rows.FieldDescriptions()))

	if nil != err {
		panic(err)
	}

	for rows.Next() {
		var cur entities.Currency

		entities.Unmarshal(rows, &cur)

		fmt.Printf("%s\n", cur.Name)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
