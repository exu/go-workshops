package main

import "fmt"
import "github.com/jmoiron/sqlx"

// import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
	db := sqlx.MustConnect("mysql", "root:root@tcp(127.0.0.1:3306)/sqlx_test")

	rows, err := db.Query("SELECT id, name, description FROM items")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		var description string

		err = rows.Scan(&id, &name, &description)

		if err != nil {
			panic(err)
		}

		fmt.Println(id, name, description)
	}
}
