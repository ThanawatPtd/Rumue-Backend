package rest

import (
	"log"

	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/exceptions"
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

	for _, user:= range *list {
		responsedUser := responses.UserDefaultResponse{}
		if err := utils.MappingParser(&user, &responsedUser); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Can't Map response",
				"log": err.Error(),
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
	id := c.Params("id")

	user, err := uh.service.GetUserByID(c.Context(), id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"log":     err.Error(),
		})
	}

	responsedUser := responses.UserDefaultResponse{}
	if err := utils.MappingParser(user, &responsedUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Can't Map response",
			"log": err.Error(),
		})
	}	

	return c.JSON(fiber.Map{
		"message": "Successful get user",
		"payload": fiber.Map{
			"user":  responsedUser,
		},
	})
}

func (uh *UserRestHandler) Register(c *fiber.Ctx) error {
	var req requests.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		log.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}
	
	createPayload := entities.User{} 
	if err := utils.MappingParser(&req, &createPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Can't map payload",
			"log": err.Error(),
		})
	}	

	selectedUser, err := uh.service.Register(c.Context(), &createPayload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	responsedUser := responses.UserRegisterResponse{}
	if err := utils.MappingParser(selectedUser, &responsedUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Can't Map response",
			"log": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": responsedUser, 
	})
}

func (uh *UserRestHandler) Login(c *fiber.Ctx) error {

	// Parse request
	var req requests.UserLoginRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	createPayload := entities.User{} 
	if err := utils.MappingParser(&req, &createPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Can't map payload",
			"log": err.Error(),
		})
	}

	// Login user
	user, token, err := uh.service.Login(c.Context(), &createPayload)
	if err != nil {
		switch err {
		case exceptions.ErrLoginFailed:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Login failed",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	responsedUser := responses.UserLoginResponse{}
	if err := utils.MappingParser(user, &responsedUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Can't Map response",
			"log": err.Error(),
		})
	}
	responsedUser.Token = token

	return c.Status(fiber.StatusOK).JSON(responsedUser)
}

func (uh *UserRestHandler) DeleteByID(c *fiber.Ctx) error {
	userID := utils.GetUserIDFromJWT(c)

	if err := uh.service.DeleteByID(c.Context(), userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)

}

func (uh *UserRestHandler) UpdateUser(c *fiber.Ctx) error {
	req := &requests.UpdateUserRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":     err,
		})
	}
	userId := utils.GetUserIDFromJWT(c)
	payload := &entities.User{}
	if err := utils.MappingParser(req, payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	updatedUser, err := uh.service.UpdateUser(c.Context(), userId, payload)	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})	
	}
	responsedUser := &responses.UserDefaultResponse{}
	if err := utils.MappingParser(updatedUser, responsedUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(responsedUser)
}
