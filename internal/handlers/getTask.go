package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/internal/models"
	"todo-api/internal/repository/dbtest"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	queryId := query.Get("id")
	queryDone := query.Get("done")

	if queryId != "" {
		id, err := strconv.Atoi(queryId)
		if err != nil {
			models.SendError(w, "invalid id", http.StatusBadRequest)
			return
		}

		for _, t := range dbtest.Tasks {
			if t.Id == id {
				if err = json.NewEncoder(w).Encode(t); err != nil {
					models.SendError(w, "internal  error", http.StatusInternalServerError)
				}
				return
			}
		}

		models.SendError(w, "task not found", http.StatusNotFound)
		return
	}

	if queryDone != "" {
		completTask, err := strconv.ParseBool(queryDone)
		if err != nil {
			models.SendError(w, "bad request", http.StatusBadRequest)
			return
		}

		var doneTasks = []*models.Task{}
		for _, t := range dbtest.Tasks {
			if t.Done == completTask {
				doneTasks = append(doneTasks, t)
			}
		}

		if err := json.NewEncoder(w).Encode(doneTasks); err != nil {
			models.SendError(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(dbtest.Tasks); err != nil {
		models.SendError(w, "server error", http.StatusInternalServerError)
	}
}
