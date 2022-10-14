package todo

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	GetTodos(ctx context.Context, filter QueryFilter) ([]Todo, error)
	GetTodo(ctx context.Context, id int) (*Todo, error)
	CreateTodo(ctx context.Context, todo *NewTodo) (*Todo, error)
	UpdateTodo(ctx context.Context, id int, todo *UpdateTodo) (*Todo, error)
	DeleteTodo(ctx context.Context, id int) (int, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) GetTodos(ctx context.Context, filter QueryFilter) ([]Todo, error) {
	var todos []Todo
	query := r.db.WithContext(ctx)

	if filter.ActivityGroupId != nil {
		if err := query.Where("activity_group_id = ?", *filter.ActivityGroupId).Find(&todos).Error; err != nil {
			return nil, err
		}

		return todos, nil
	}

	if err := query.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *repository) GetTodo(ctx context.Context, id int) (*Todo, error) {
	var todo Todo

	if err := r.db.WithContext(ctx).First(&todo, id).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *repository) CreateTodo(ctx context.Context, newTodo *NewTodo) (*Todo, error) {
	todo := Todo{
		ActivityGroupId: newTodo.ActivityGroupId,
		Title:           newTodo.Title,
	}

	if err := r.db.WithContext(ctx).Create(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *repository) UpdateTodo(ctx context.Context, id int, updateTodo *UpdateTodo) (*Todo, error) {
	todo := Todo{}
	if err := r.db.WithContext(ctx).First(&todo, id).Error; err != nil {
		return nil, err
	}

	todo.UpdatedAt = time.Now()
	todo.Title = updateTodo.Title

	if err := r.db.WithContext(ctx).Save(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *repository) DeleteTodo(ctx context.Context, id int) (int, error) {
	todo := Todo{}

	if err := r.db.WithContext(ctx).First(&todo, id).Error; err != nil {
		return -1, err
	}

	if err := r.db.WithContext(ctx).Delete(&todo, id).Error; err != nil {
		return -1, err
	}

	return todo.Id, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
