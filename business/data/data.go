// Package data adalah factory dari business repository
package data

import (
	"github.com/Budi721/todolistskyshi/business/data/activity"
	"github.com/Budi721/todolistskyshi/business/data/todo"
	"gorm.io/gorm"
)

type Factory struct {
	TodoRepository     todo.Repository
	ActivityRepository activity.Repository
}

func NewFactory(db *gorm.DB) *Factory {
	return &Factory{
		todo.NewRepository(db),
		activity.NewRepository(db),
	}
}
