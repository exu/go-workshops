package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Item struct {
	Id          int
	Name        string
	Description string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:7701)/gotraining")
	defer db.Close()

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	item := Item{}
	db.First(&item)
	fmt.Println(item)
}
