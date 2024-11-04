package rest

import (
	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/exceptions"
	"github.com/ThanawatPtd/SAProject/domain/requests"
	"github.com/ThanawatPtd/SAProject/domain/responses"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService usecases.AuthUseCase
}

func ProvideAuthRestHandler(authService usecases.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (ah *AuthHandler) Register(c *fiber.Ctx) error {
	req := requests.CreateUserRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := utils.ValidateStruct(req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	createPayload := entities.User{}
	if err := utils.MappingParser(&req, &createPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := ah.authService.Register(c.Context(), &createPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create Success",
	})
}

func (uh *AuthHandler) Login(c *fiber.Ctx) error {
	// Parse request
	var req requests.UserLoginRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	createPayload := entities.User{}
	if err := utils.MappingParser(&req, &createPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Login user
	token, err := uh.authService.Login(c.Context(), &createPayload)
	if err != nil {
		switch err {
		case exceptions.ErrLoginFailed:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": exceptions.ErrLoginFailed.Error(),
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	responsedUser := responses.UserLoginResponse{}
	responsedUser.Token = token

	return c.Status(fiber.StatusOK).JSON(responsedUser)
}
