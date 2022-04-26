package models

import "gorm.io/gorm"
type Cart struct {
	gorm.Model
	Id uint 
	AccountRefer uint 
	CartItems []CartItem `gorm:"foreignKey:CartRefer"`
}
