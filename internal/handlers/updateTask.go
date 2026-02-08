package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/internal/models"
	"todo-api/internal/repository/dbtest"
)

type DoneTask struct {
	Done bool `json:"done"`
}

type UpTask struct {
	Updated bool `json:"updated"`
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	queryId := query.Get("id")

	if queryId == "" {
		models.SendError(w, "task not found", http.StatusNotFound)
		return
	}

	var done DoneTask

	err := json.NewDecoder(r.Body).Decode(&done)
	if err != nil {
		models.SendError(w, "bad request", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(queryId)
	if err != nil {
		models.SendError(w, "invalid id", http.StatusBadRequest)
		return
	}

	var isFound bool
	for _, t := range dbtest.Tasks {
		if t.Id == id {
			t.Done = done.Done
			isFound = true
			break
		}
	}

	if !isFound {
		models.SendError(w, "task not found", http.StatusNotFound)
		return
	}

	var update UpTask
	update.Updated = true

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(update)
	if err != nil {
		models.SendError(w, "status server error", http.StatusInternalServerError)
		return
	}
}
