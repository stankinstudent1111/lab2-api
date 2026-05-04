package main

import (
	"log"
	"net/http"

	"lab2-api/internal/handlers"
	"lab2-api/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handlers.PingHandler)
	wrappedMux := middleware.Logging(mux)
	log.Printf("Сервер запущен на порту %s", ":8080")

	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		log.Fatal(err)
	}
}
