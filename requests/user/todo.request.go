package userrequests

type TodoRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

type UpdateTodoRequest struct {
	TodoRequest
	ID string `param:"id" validate:"required,uuid"`
}
