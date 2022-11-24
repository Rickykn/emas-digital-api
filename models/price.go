package models

type Price struct {
	ID            int    `json:"id" gorm:"primary_key"`
	Topup_price   int    `json:"topup_price"`
	Buyback_price int    `json:"buyback_price"`
	Admin_id      string `json:"admin_id"`
}

type PriceDTO struct {
	Topup_price   int    `json:"topup_price"`
	Buyback_price int    `json:"buyback_price"`
	Admin_id      string `json:"admin_id"`
}
