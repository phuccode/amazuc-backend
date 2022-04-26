package database

import (
	"../models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "sqlserver://sa:dinhphuc98@LAPTOP-ODRLSC22:1433?database=authGolang"
	connection, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.ProductStatus{},
		&models.Category{},
		&models.Product{},
		&models.ProductImage{},
		&models.Cart{},
		&models.CartItem{},
		&models.OrderBill{},
	)
}
