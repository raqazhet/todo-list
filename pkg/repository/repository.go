package repository

type (
	Authorization interface{}
	TodoList      interface{}
	TodoItem      interface{}
)

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{}
}
