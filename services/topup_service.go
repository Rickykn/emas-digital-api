package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rickykn/emas-digital-api/helpers"
	"github.com/Rickykn/emas-digital-api/models"
)

func Topup(w http.ResponseWriter, r *http.Request) *models.TopupDTO {
	body := &models.TopupDTO{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(body)
	if err != nil {
		helpers.WriteResponse(
			w,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		)
		return nil
	}
	fmt.Println(body)

	return body
}
