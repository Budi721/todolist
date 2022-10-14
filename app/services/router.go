package services

import (
	activityapi "github.com/Budi721/todolistskyshi/app/services/activity-api"
	todoapi "github.com/Budi721/todolistskyshi/app/services/todo-api"
	"github.com/Budi721/todolistskyshi/business/data"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewAppRouter(app *fiber.App, factory *data.Factory, validate *validator.Validate) {
	activityapi.NewRouter(app, factory, validate)
	todoapi.NewRouter(app, factory, validate)
}
