package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Id       uint
	Name     string      `json:"name"`
	Email    string      `json:"email" gorm:"unique"`
	Password []byte      `json:"-"`
	Orders   []OrderBill `gorm:"foreignKey:AccountRefer"`
	UserID   uint
	Carts    []Cart `gorm:"foreignKey:AccountRefer"`
}
