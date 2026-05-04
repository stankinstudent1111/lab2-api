package main

import (
	"log"
	"net/http"

	"lab2-api/internal/handlers"
	"lab2-api/internal/middleware"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handlers.PingHandler)
	wrappedMux := middleware.Logging(mux)
	log.Printf("Сервер запущен на порту %s", port)

	err := http.ListenAndServe(port, wrappedMux)
	if err != nil {
		log.Fatal(err)
	}
}
