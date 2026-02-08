package handlers

import (
	"encoding/json"
	"net/http"
	"todo-api/internal/models"
)

type ExternalTodo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetExternalTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		models.SendError(w, "failed to fetch external data", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	var externalTasks []ExternalTodo
	if err := json.NewDecoder(resp.Body).Decode(&externalTasks); err != nil {
		models.SendError(w, "failed to parse external json", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(externalTasks); err != nil {
		models.SendError(w, "internal server error", http.StatusInternalServerError)
	}
}
