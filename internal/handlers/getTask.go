package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	models2 "todo-api/internal/models"
	"todo-api/internal/repository/dbtest"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	queryId := query.Get("id")

	if queryId == "" {
		if err := json.NewEncoder(w).Encode(dbtest.Tasks); err != nil {
			models2.SendError(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	id, err := strconv.Atoi(queryId)
	if err != nil {
		models2.SendError(w, "invalid id", http.StatusBadRequest)
		return
	}

	var foundTask *models2.Task

	for _, t := range dbtest.Tasks {
		if t.Id == id {
			foundTask = t
			break
		}
	}

	if foundTask == nil {
		models2.SendError(w, "task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(foundTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
