package main

import "fmt"
import "database/sql"
import "github.com/jmoiron/sqlx"

// We're using MySQL
import _ "github.com/go-sql-driver/mysql"

// We can map columns to fields
type Item struct {
	Id  int            `db:"id"`
	Nme sql.NullString `db:"name"`
	Dsc sql.NullString `db:"description"`
}

func main() {

	db := sqlx.MustConnect("mysql", "root:root@tcp(127.0.0.1:3306)/sqlx_test")
	rows, err := db.Queryx("SELECT id, name, description FROM items")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var item Item
		err = rows.StructScan(&item)

		if err != nil {
			panic(err)
		}

		fmt.Printf(
			"%d - %s:  %s\n===================\n",
			item.Id,
			item.Nme.String,
			item.Dsc.String,
		)
	}
}
