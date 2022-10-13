package activity_api

import (
	"github.com/Budi721/todolistskyshi/business/data"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, factory *data.Factory, validate *validator.Validate) {
	handler := NewHandler(factory.ActivityRepository, validate)
	todoRoute := app.Group("/activity-groups")
	todoRoute.Get("", handler.GetActivitiesHandler)
	todoRoute.Get("/:id", handler.GetActivityHandler)
	todoRoute.Post("", handler.PostActivityHandler)
	todoRoute.Delete("/:id", handler.DeleteActivityHandler)
	todoRoute.Patch("/:id", handler.PatchActivityHandler)
}
