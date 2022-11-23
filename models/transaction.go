package models

import "time"

type Transaction struct {
	ID         int     `json:"id" gorm:"primary_key"`
	Price_id   *int    `json:"price_id"`
	Price      Price   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Account_id *int    `json:"account_id"`
	Account    Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Gram       float32 `json:"gram"`
	Type       string  `json:"type"`
	CreatedAt  time.Time
}
