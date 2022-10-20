package todo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	Id              uuid.UUID      `gorm:"primaryKey" gorm:"type:uuid"`
	ActivityGroupId int            `json:"activity_group_id"`
	Title           string         `json:"title"`
	IsActive        string         `json:"is_active"  gorm:"default:1"`
	Priority        string         `json:"priority"  gorm:"default:very-high"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	t.Id = uuid.New()
	return
}

type QueryFilter struct {
	ActivityGroupId *int
}

type NewTodo struct {
	ActivityGroupId int    `json:"activity_group_id"  validate:"required"`
	Title           string `json:"title"  validate:"required"`
}

type UpdateTodo struct {
	Title string `json:"title"  validate:"required"`
}
