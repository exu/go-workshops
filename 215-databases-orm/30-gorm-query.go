package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id          int64
	Name        string
	Description string
	Age         float64
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:7701)/gotraining?charset=utf8")
	defer db.Close()

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	// ;	db.AutoMigrate(&User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	users := []User{
		{Id: 1, Name: "Wacek", Description: "Kowal", Age: 23.123},
		{Id: 2, Name: "Micek", Description: "Programista", Age: 18.123},
		{Id: 3, Name: "Pucek", Description: "Spawacz", Age: 54.123},
		{Id: 4, Name: "Gigus", Description: "Przedszkolak", Age: 4.123},
		{Id: 5, Name: "Migus", Description: "Nauczyciel", Age: 34.123},
		{Id: 6, Name: "Figus", Description: "Przedszkolak", Age: 4.392},
	}

	for _, user := range users {
		db.FirstOrCreate(&user)
	}

	// find all
	result := []User{}
	db.Find(&result)
	fmt.Println(result)

	user := User{}
	db.Where("name = ?", "Pucek").First(&user)
	fmt.Println("One user = Pucek", user)

	db.Where("name = ?", "Pucek").Find(&users)
	fmt.Println("name = Pucek", users)

	db.Where("name <> ?", "Pucek").Find(&users)
	fmt.Println("name <> Pucek", users)

	// IN
	db.Where("name in (?)", []string{"Pucek", "Wacek"}).Find(&users)
	fmt.Println("name in (Pucek Wacek)", users)

	// LIKE
	db.Where("name LIKE ?", "Wa%").Find(&users)
	fmt.Println("name LIKE ek", users)

	// AND
	db.Where("name = ? and description = ?", "Figus", "Przedszkolak").Find(&users)
	fmt.Println("name = Figus and descr = Przedszkolak", users)
}
