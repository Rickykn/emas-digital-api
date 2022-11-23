package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type jsonResponse struct {
	Code    string
	Message string
	Data    interface{}
}

func WriteResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	code := strings.Replace(strings.ToUpper(http.StatusText(statusCode)), " ", "_", -1)
	response := &jsonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}
