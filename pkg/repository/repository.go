package repository

type Authorization interface {
}

type UsersList interface {
}

type Repository struct {
	Authorization
	UsersList
}

func NewRepository() *Repository {
	return &Repository{}
}
