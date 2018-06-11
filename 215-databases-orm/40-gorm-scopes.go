package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Order struct {
	Id          int
	Name        string
	Description string
}

func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
	return db.Where("amount > ?", 1000)
}

func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
	return db.Where("pay_mode_sign = ?", "C")
}

func PaidWithCod(db *gorm.DB) *gorm.DB {
	return db.Where("pay_mode_sign = ?", "C")
}

func OrderStatus(status []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
	}
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:7701)/gotraining")
	defer db.Close()

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	orders := []Order{}

	db.Scopes(AmountGreaterThan1000, PaidWithCreditCard).Find(&orders)
	// Find all credit card orders and amount greater than 1000

	db.Scopes(AmountGreaterThan1000, PaidWithCod).Find(&orders)
	// Find all COD orders and amount greater than 1000

	db.Scopes(OrderStatus([]string{"paid", "shipped"})).Find(&orders)
	// Find all paid, shipped orders
}
