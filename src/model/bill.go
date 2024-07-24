package model

type bill struct {
	Id       string    `json:"id"`
	BillType bill_type `json:"bill_type"`
}
