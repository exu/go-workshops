package main

import "fmt"
import "database/sql"
import "github.com/jmoiron/sqlx"

// We're using MySQL
import _ "github.com/go-sql-driver/mysql"

type Item struct {
	Id          int            `db:"id"`
	Name        sql.NullString `db:"name"`
	Description sql.NullString `db:"description"`
}

func main() {
	var item Item

	// Must.... functions will panic on fail
	db := sqlx.MustConnect("mysql", "root:root@tcp(127.0.0.1:3306)/sqlx_test")

	row := db.QueryRowx("SELECT * FROM items WHERE id=?", 3)
	err := row.StructScan(&item)

	if err != nil {
		panic(err)
	}

	fmt.Println(item)
}
