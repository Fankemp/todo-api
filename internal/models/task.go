package models

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func NewTask(id int, title string, done bool) *Task {
	return &Task{
		Id:    id,
		Title: title,
		Done:  done,
	}
}
