package main

import (
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name        string
	Description string
	Counters    float64
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:7701)/gotraining")
	defer db.Close()

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	// Create table
	// db.CreateTable(&User{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})

	// drop table
	// db.DropTable(&User{})

	// ModifyColumn
	// db.Model(&User{}).ModifyColumn("description", "text")

	// DropColumn
	// db.Model(&User{}).DropColumn("description")

	// Automating Migration
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	// db.AutoMigrate(&User{}, &Product{}, &Order{})
}
