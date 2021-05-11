package entities

type Task struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
