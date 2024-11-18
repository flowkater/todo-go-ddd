package command

type CreateTodoCommand struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodoCommand struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ToggleTodoCommand struct {
	ID int `json:"id"`
}

type DeleteTodoCommand struct {
	ID int `json:"id"`
}
