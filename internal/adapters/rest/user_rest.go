package rest

import (
	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/requests"
	"github.com/ThanawatPtd/SAProject/domain/responses"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/gofiber/fiber/v2"
)

type UserRestHandler struct {
	service usecases.UserUseCase
}

func ProvideUserRestHandler(userUseCase usecases.UserUseCase) *UserRestHandler {
	return &UserRestHandler{service: userUseCase}
}

func (uh *UserRestHandler) GetUsers(c *fiber.Ctx) error {
	list, err := uh.service.GetUsers(c.Context())

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"log":     err.Error(),
		})
	}

	var responsedUsers []responses.UserDefaultResponse

	for _, user := range list {
		responsedUser := responses.UserDefaultResponse{}
		if err := utils.MappingParser(&user, &responsedUser); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Can't Map response",
				"log":     err.Error(),
			})
		}
		responsedUsers = append(responsedUsers, responsedUser)
	}

	return c.JSON(fiber.Map{
		"message": "Successful get all users",
		"payload": fiber.Map{
			"user": responsedUsers,
		},
	})
}

func (uh *UserRestHandler) GetUserByID(c *fiber.Ctx) error {
	jwt := utils.GetJWTFromContext(c)
	user, err := uh.service.GetUserByID(c.Context(), jwt.UserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":     err.Error(),
		})
	}

	responsedUser := responses.UserDefaultResponse{}
	if err := utils.MappingParser(user, &responsedUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":     err.Error(),
		})
	}

	return c.JSON(responsedUser)
}

func (uh *UserRestHandler) DeleteByID(c *fiber.Ctx) error {
	jwt := utils.GetJWTFromContext(c)
	if err := uh.service.DeleteByID(c.Context(), jwt.UserID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":     err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)

}

func (uh *UserRestHandler) UpdateUser(c *fiber.Ctx) error {
	req := &requests.UpdateUserRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := utils.ValidateStruct(*req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	payload := &entities.User{}
	if err := utils.MappingParser(req, payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	jwt := utils.GetJWTFromContext(c)
	payload.ID = jwt.UserID
	updatedUser, err := uh.service.UpdateUser(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	responsedUser := &responses.UserDefaultResponse{}
	if err := utils.MappingParser(updatedUser, responsedUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(responsedUser)
}

func (uh *UserRestHandler) UpdatePassword(c *fiber.Ctx) error {
	req := &requests.UpdatePasswordRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := utils.ValidateStruct(*req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	jwt := utils.GetJWTFromContext(c)
	if err := uh.service.UpdatePassword(c.Context(), jwt.UserID, req.OldPassword, req.NewPassword); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
