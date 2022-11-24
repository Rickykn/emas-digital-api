package models

type Account struct {
	ID             int `json:"id" gorm:"primary_key"`
	Number         int `json:"number"`
	Account_amount int `json:"Account_amount"`
}

type TopupDTO struct {
	Gram           float32 `json:"gram"`
	Price          int     `json:"price"`
	Account_amount int     `json:"Account_amount"`
}
