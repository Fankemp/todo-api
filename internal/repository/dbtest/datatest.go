package dbtest

import (
	"sync"
	"todo-api/internal/models"
)

var (
	Tasks = []*models.Task{}
	Mu    sync.Mutex
)
