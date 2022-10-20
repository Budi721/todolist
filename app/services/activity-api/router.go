package activity_api

import (
	"github.com/Budi721/todolistskyshi/business/data"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, factory *data.Factory, validate *validator.Validate, translator ut.Translator) {
	handler := NewHandler(factory.ActivityRepository, validate, translator)
	todoRoute := app.Group("/activity-groups")
	todoRoute.Get("", handler.GetActivitiesHandler)
	todoRoute.Get("/:id", handler.GetActivityHandler)
	todoRoute.Post("", handler.PostActivityHandler)
	todoRoute.Delete("/:id", handler.DeleteActivityHandler)
	todoRoute.Patch("/:id", handler.PatchActivityHandler)
}
