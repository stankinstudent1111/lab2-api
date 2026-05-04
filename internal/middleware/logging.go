package middleware

import (
	"log"
	"net/http"
)

// Middleware для логирования запросов
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		log.Printf("Запрос пришел")
	})
}
