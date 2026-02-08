package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/internal/models"
	"todo-api/internal/repository/dbtest"
)

type DeleteDone struct {
	Message string `json:"message"`
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	queryId := query.Get("id")

	id, err := strconv.Atoi(queryId)
	if err != nil {
		models.SendError(w, "invalid id", http.StatusBadRequest)
		return
	}

	index := -1
	for i, t := range dbtest.Tasks {
		if t.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		models.SendError(w, "task not found", http.StatusNotFound)
		return
	}

	dbtest.Tasks = append(dbtest.Tasks[:index], dbtest.Tasks[index+1:]...)

	var isDeleted DeleteDone
	isDeleted.Message = "task deleted"
	if err = json.NewEncoder(w).Encode(isDeleted); err != nil {
		models.SendError(w, "internal server error", http.StatusInternalServerError)
	}
}
