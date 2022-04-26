package models

type ProductStatus struct {
	Id       uint
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignKey:ProductStRefer"`
}
