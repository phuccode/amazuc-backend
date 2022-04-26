package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id             uint
	Name           string `json:"name"`
	Price          string `json:"price"`
	Quantity       string `json:"quatity"`
	Description    string `json:"description"`
	CategoryRefer  uint
	ProductStRefer uint
	ProductImages  []ProductImage `gorm:"foreignKey:ProductRefer"`
	OrderBills     []OrderBill    `gorm:"foreignKey:ProductRefer"`
	CartItems      []CartItem     `gorm:"foreignKey:ProductRefer"`
}
