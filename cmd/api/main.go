package main

import (
	"log"
	"net/http"

	"lab2-api/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handlers.PingHandler)
	log.Printf("Сервер запущен на порту %s", ":8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
