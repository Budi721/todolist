package todo

import "gorm.io/gorm"

type Repository interface {
	GetTodos()
	GetTodo(id uint)
	//CreateBook(id uint, title string, author string, price uint)
	//UpdateBook(id uint, title string, author string, price uint)
	//DeleteBook(id uint) (*model.Book, error)
}

type repository struct {
	db *gorm.DB
}

func (r repository) GetTodos() {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetTodo(id uint) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
