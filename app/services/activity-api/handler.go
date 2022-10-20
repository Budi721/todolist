package activity_api

import (
	"errors"
	"fmt"
	"github.com/Budi721/todolistskyshi/business/data/activity"
	"github.com/Budi721/todolistskyshi/fondation/web"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

type Handler struct {
	repository activity.Repository
	validate   *validator.Validate
	trans      ut.Translator
}

func NewHandler(repository activity.Repository, validate *validator.Validate, trans ut.Translator) *Handler {
	return &Handler{repository: repository, validate: validate, trans: trans}
}

func (h *Handler) GetActivitiesHandler(c *fiber.Ctx) error {
	activities, err := h.repository.GetActivities(c.Context())

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.Status(fiber.StatusNotFound).JSON(web.Response{
				Status:  "Not Found",
				Message: "Not Found activities",
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
		Data:    activities,
	})
}

func (h *Handler) GetActivityHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	result, err := h.repository.GetActivity(c.Context(), id)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.Status(fiber.StatusNotFound).JSON(web.Response{
				Status:  "Not Found",
				Message: fmt.Sprintf("Activity with ID %v Not Found", id),
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

func (h *Handler) PostActivityHandler(c *fiber.Ctx) error {
	payload := new(activity.NewActivity)
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	err := h.validate.Struct(payload)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			Status:  "Bad Request",
			Message: errs[0].Translate(h.trans),
			Data:    nil,
		})
	}

	result, err := h.repository.CreateActivity(c.Context(), payload)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(web.Response{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	})
}

func (h *Handler) PatchActivityHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			Status:  "Bad Request",
			Message: "cannot parse id",
			Data:    nil,
		})
	}

	payload := new(activity.UpdateActivity)
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	err = h.validate.Struct(payload)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			Status:  "Bad Request",
			Message: errs[0].Translate(h.trans),
			Data:    nil,
		})
	}

	result, err := h.repository.UpdateActivity(c.Context(), id, payload)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.Status(fiber.StatusNotFound).JSON(web.Response{
				Status:  "Not Found",
				Message: fmt.Sprintf("Activity with ID %v Not Found", id),
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

func (h *Handler) DeleteActivityHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			Status:  "Bad Request",
			Message: "Bad Request",
			Data:    nil,
		})
	}

	_, err = h.repository.DeleteActivity(c.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.Status(fiber.StatusNotFound).JSON(web.Response{
				Status:  "Not Found",
				Message: fmt.Sprintf("Activity with ID %v Not Found", id),
				Data:    nil,
			})
		default:
			return err
		}
	}

	return c.JSON(web.Response{
		Status:  "Success",
		Message: "Success",
		Data:    struct{}{},
	})
}
