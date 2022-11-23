package handlers

import (
	"net/http"

	"github.com/Rickykn/emas-digital-api/helpers"
)

func Home(w http.ResponseWriter, r *http.Request) {
	helpers.WriteResponse(w, http.StatusOK, "Hello world!", nil)
}
