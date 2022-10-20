package services

import (
	activityapi "github.com/Budi721/todolistskyshi/app/services/activity-api"
	todoapi "github.com/Budi721/todolistskyshi/app/services/todo-api"
	"github.com/Budi721/todolistskyshi/business/data"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewAppRouter(app *fiber.App, factory *data.Factory, validate *validator.Validate, translator ut.Translator) {
	activityapi.NewRouter(app, factory, validate, translator)
	todoapi.NewRouter(app, factory, validate, translator)
}
