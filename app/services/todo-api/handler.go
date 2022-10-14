package todo_api

import (
	"errors"
	"fmt"
	"github.com/Budi721/todolistskyshi/business/data/todo"
	"github.com/Budi721/todolistskyshi/fondation/web"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

type Handler struct {
	repository todo.Repository
	validate   *validator.Validate
}

func NewHandler(repository todo.Repository, validate *validator.Validate) *Handler {
	return &Handler{repository: repository, validate: validate}
}

func (h *Handler) GetTodosHandler(c *fiber.Ctx) error {
	var id *int
	fmt.Println(c.Query("activity_group_id") != "")
	if c.Query("activity_group_id") != "" {
		convert, _ := strconv.Atoi(c.Query("activity_group_id"))
		id = &convert
	}

	todos, err := h.repository.GetTodos(c.Context(), todo.QueryFilter{ActivityGroupId: id})

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.Status(fiber.StatusNotFound).JSON(web.Response{
				Status:  "Not found",
				Message: "Not found",
				Data:    nil,
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(web.Response{
				Status:  "Internal server error",
				Message: "Internal server error",
				Data:    nil,
			})
		}
	}

	return c.JSON(web.Response{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

func (h *Handler) GetTodoHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	result, err := h.repository.GetTodo(c.Context(), id)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.Status(fiber.StatusNotFound).JSON(web.Response{
				Status:  "Not found",
				Message: "Not found",
				Data:    nil,
			})
		default:
			return err
		}
	}

	return c.JSON(web.Response{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	})
}

func (h *Handler) PostTodoHandler(c *fiber.Ctx) error {
	payload := new(todo.NewTodo)
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	err := h.validate.Struct(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			Status:  "Bad request",
			Message: "Bad request",
			Data:    nil,
		})
	}

	result, err := h.repository.CreateTodo(c.Context(), payload)
	if err != nil {
		return err
	}

	return c.JSON(web.Response{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	})
}

func (h *Handler) PatchTodoHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			Status:  "Bad request",
			Message: "Bad request",
			Data:    nil,
		})
	}

	payload := new(todo.UpdateTodo)
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	err = h.validate.Struct(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			Status:  "Bad request",
			Message: "Bad request",
			Data:    nil,
		})
	}

	result, err := h.repository.UpdateTodo(c.Context(), id, payload)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.Status(fiber.StatusNotFound).JSON(web.Response{
				Status:  "Not found",
				Message: "Not found",
				Data:    nil,
			})
		default:
			return err
		}
	}

	return c.JSON(web.Response{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	})
}

func (h *Handler) DeleteTodoHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			Status:  "Bad request",
			Message: "Bad request",
			Data:    nil,
		})
	}

	result, err := h.repository.DeleteTodo(c.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.Status(fiber.StatusNotFound).JSON(web.Response{
				Status:  "Not found",
				Message: "Not found",
				Data:    nil,
			})
		default:
			return err
		}
	}

	return c.JSON(web.Response{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	})
}
