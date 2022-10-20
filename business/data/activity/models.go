package activity

import (
	"github.com/Budi721/todolistskyshi/business/data/todo"
	"gorm.io/gorm"
	"time"
)

type Activity struct {
	Id        int            `gorm:"primaryKey" gorm:"column:activity_group_id" json:"id"`
	Email     string         `json:"email"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Todos     []todo.Todo    `gorm:"foreignKey:ActivityGroupId" json:"-"`
}

type NewActivity struct {
	Title string `json:"title"  validate:"required"`
	Email string `json:"email"  validate:"required"`
}

type UpdateActivity struct {
	Title string `json:"title" validate:"required"`
}
