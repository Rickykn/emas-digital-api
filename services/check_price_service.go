package services

import (
	"github.com/Rickykn/emas-digital-api/database"
	"github.com/Rickykn/emas-digital-api/models"
)

func CheckPrice() *models.Price {
	DB := database.Get()
	var price *models.Price
	result := DB.Last(&price)

	if result.Error != nil {
		panic(result.Error)
	}
	return price
}
