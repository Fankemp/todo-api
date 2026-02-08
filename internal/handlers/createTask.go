package handlers

import (
	"encoding/json"
	"net/http"
	models2 "todo-api/internal/models"
	"todo-api/internal/repository/dbtest"
)

type TitleTask struct {
	Text string `json:"title"`
}

var id = 0

func NextId() int {
	id++
	return id
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbtest.Mu.Lock()
	defer dbtest.Mu.Unlock()

	var text TitleTask
	err := json.NewDecoder(r.Body).Decode(&text)
	if err != nil {
		models2.SendError(w, "bad request", http.StatusBadRequest)
		return
	}

	if text.Text == "" {
		models2.SendError(w, "invalid title", http.StatusBadRequest)
		return
	}

	var newTask *models2.Task

	newTask = models2.NewTask(NextId(), text.Text, false)
	dbtest.Tasks = append(dbtest.Tasks, newTask)

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(newTask); err != nil {
		models2.SendError(w, "internal server error", http.StatusInternalServerError)
		return
	}

}
