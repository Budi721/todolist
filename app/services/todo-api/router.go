package todo_api

import (
	"github.com/Budi721/todolistskyshi/business/data"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, factory *data.Factory, validate *validator.Validate) {
	handler := NewHandler(factory.TodoRepository, validate)
	todoRoute := app.Group("/todo-items")
	todoRoute.Get("", handler.GetTodosHandler)
	todoRoute.Get("/:id", handler.GetTodoHandler)
	todoRoute.Post("", handler.PostTodoHandler)
	todoRoute.Delete("/:id", handler.DeleteTodoHandler)
	todoRoute.Patch("/:id", handler.PatchTodoHandler)
}
