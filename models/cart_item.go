package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Quantity     string `json:"quantity"`
	CartRefer    uint
	ProductRefer uint
}
