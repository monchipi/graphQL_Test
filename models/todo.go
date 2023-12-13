package models

type NewTodo struct {
	Text   string `json:"Text"`
	UserID string `json:"UserId"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
  IsCompleted bool   `json:"IsCompleted"`
}

