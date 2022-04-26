package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint
	Name     string  `json:"name"`
	Birthday string  `json:"birthday"`
	Phone    string  `json:"phone"`
	Avatar   string  `json:"avatar"`
	Address  string  `json:"address"`
	Accounts Account `gorm:"foreignKey: UserID"`
}
