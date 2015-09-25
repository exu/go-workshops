package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// CREATE DATABASE gotraining;
// CREATE TABLE squares (number int, squareNumber int);
func main() {
	db, err := sql.Open("mysql", "root:root@tcp(l:3306)/gotraining")
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

	// czytanie
	stmtOut, err := db.Prepare("SELECT number, squareNumber FROM squares WHERE number = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var num, squareNum int
	err = stmtOut.QueryRow(13).Scan(&num, &squareNum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 13 is: %d\n", squareNum)
}
