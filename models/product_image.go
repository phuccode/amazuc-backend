package models

type ProductImage struct {
	Id           uint
	Image        string `json:"image"`
	ProductRefer uint
}
