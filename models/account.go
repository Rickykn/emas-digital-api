package models

type Account struct {
	ID             int `json:"id" gorm:"primary_key"`
	Number         int `json:"number"`
	Account_amount int `json:"Account_amount"`
}
