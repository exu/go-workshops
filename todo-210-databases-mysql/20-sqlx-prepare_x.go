package main

import "fmt"
import "github.com/jmoiron/sqlx"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Item struct {
	Id          int            `db:"id"`
	Name        sql.NullString `db:"name"`
	Description sql.NullString `db:"description"`
}

func main() {
	var item Item

	db := sqlx.MustConnect("mysql", "root:root@tcp(127.0.0.1:3306)/sqlx_test")
	stmt, err := db.Preparex(`SELECT * FROM items WHERE id=?`)

	// existing one
	err = stmt.Get(&item, 1)
	fmt.Println(item)

	// not existing one
	err = stmt.Get(&item, 3)
	fmt.Println(item)

	// handling non existing item
	if err == sql.ErrNoRows {
		fmt.Println("There is no row with id", 900)
	} else if err != nil {
		panic(err)
	}
}
