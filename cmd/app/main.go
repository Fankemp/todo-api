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

	midd := middleware.ApiKeyMiddleware(http.DefaultServeMux)

	if err := http.ListenAndServe(":8080", midd); err != nil {
		log.Fatal(err)
	}
}
