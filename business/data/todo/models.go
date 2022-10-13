package todo

import "time"

type Todo struct {
	Id              int       `json:"id"`
	ActivityGroupId int       `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        string    `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}
