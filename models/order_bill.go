package models

type OrderBill struct {
	Id           string
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Total        string `json:"total"`
	MethodPay    string `json:"method_pay"`
	AccountRefer uint
	ProductRefer uint
}
