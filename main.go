package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rickykn/emas-digital-api/database"
	"github.com/Rickykn/emas-digital-api/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	errConnect := database.Connect()

	if errConnect != nil {
		panic(errConnect)
	}
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Home).Methods(http.MethodGet)

	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}
