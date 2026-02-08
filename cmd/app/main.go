package main

import (
	"log"
	"net/http"
	"todo-api/internal/handlers"
	"todo-api/internal/middleware"
)

func main() {
	http.HandleFunc("GET /tasks", handlers.GetTask)
	http.HandleFunc("POST /tasks", handlers.CreateTask)
	http.HandleFunc("PATCH  /tasks", handlers.UpdateTask)

	var handler http.Handler = http.DefaultServeMux
	handler = middleware.ApiKeyMiddleware(http.DefaultServeMux)
	handler = middleware.LoggerMiddleware(handler)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
