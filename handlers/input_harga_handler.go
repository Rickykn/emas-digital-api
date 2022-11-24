package handlers

import (
	"net/http"

	"github.com/Rickykn/emas-digital-api/helpers"
	"github.com/Rickykn/emas-digital-api/services"
)

func HandlerInputHarga(w http.ResponseWriter, r *http.Request) {
	b := services.InputPrice(w, r)
	services.ReceiveFromKafka()

	helpers.WriteResponse(w, http.StatusOK, "Success add input harga", b)

}

func GetHarga(w http.ResponseWriter, r *http.Request) {
	b := services.CheckPrice()

	helpers.WriteResponse(w, http.StatusOK, "Get Success", b)

}

func Topup(w http.ResponseWriter, r *http.Request) {
	b := services.Topup(w, r)

	helpers.WriteResponse(w, http.StatusOK, "Success Topup", b)

}
