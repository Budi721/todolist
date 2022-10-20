package todo_api

import (
	"github.com/Budi721/todolistskyshi/business/data"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, factory *data.Factory, validate *validator.Validate, translator ut.Translator) {
	handler := NewHandler(factory.TodoRepository, validate, translator)
	todoRoute := app.Group("/todo-items")
	todoRoute.Get("", handler.GetTodosHandler)
	todoRoute.Get("/:id", handler.GetTodoHandler)
	todoRoute.Post("", handler.PostTodoHandler)
	todoRoute.Delete("/:id", handler.DeleteTodoHandler)
	todoRoute.Patch("/:id", handler.PatchTodoHandler)
}
