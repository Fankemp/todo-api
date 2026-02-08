package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeLog := time.Now().Format("2006-01-02T15:04:05")

		method := r.Method
		path := r.URL.Path

		message := "request received"

		fmt.Printf("%s %s %s %s\n", timeLog, method, path, message)

		next.ServeHTTP(w, r)
	})
}
