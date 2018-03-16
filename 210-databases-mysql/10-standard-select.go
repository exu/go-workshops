package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:7701)/gotraining")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO squares VALUES( ?, ? )")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec(i, (i * i))
		if err != nil {
			panic(err.Error())
		}
	}

	// red from tables
	stmtOut, err := db.Prepare("SELECT num, square FROM squares WHERE num = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var num, square int
	err = stmtOut.QueryRow(13).Scan(&num, &square)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of %d is: %d\n", num, square)
}
