package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Id       uint
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignKey:CategoryRefer"`
}
