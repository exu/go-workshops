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
	// Must.... functions will panic on fail
	db := sqlx.MustConnect("mysql", "root:root@tcp(127.0.0.1:3306)/sqlx_test")
	var item Item

	// We'll get most recent item and map it into our struct
	err := db.Get(&item, "SELECT * FROM items ORDER BY id DESC LIMIT 1")
	if err != nil {
		panic(err)
	}

	fmt.Printf("id: %d, %s, %s", item.Id, item.Name.String, item.Description.String)
}
