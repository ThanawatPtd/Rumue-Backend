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

	return c.JSON(fiber.Map{
		"message": "Successful get user",
		"payload": fiber.Map{
			"user": list,
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

	return c.JSON(fiber.Map{
		"message": "Successful get user",
		"payload": fiber.Map{
			"user": responses.UserDefaultResponse{
				ID:          user.UserId,
				Name:        user.Fname + " " + user.Lname,
				Email:       user.Email,
				PhoneNumber: user.PhoneNumber,
				Address:     user.Address,
			},
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

	createPayload := entities.User{
		UserId:      "",
		Email:       req.Email,
		Fname:       req.Fname,
		Lname:       req.Lname,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
	}
	selectedUser, err := uh.service.Register(c.Context(), &createPayload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": responses.UserRegisterResponse{
			ID:          selectedUser.UserId,
			Name:        selectedUser.Fname + selectedUser.Lname,
			Email:       selectedUser.Email,
			PhoneNumber: selectedUser.PhoneNumber,
			Address:     selectedUser.Address,
		},
	})
}

func (uh *UserRestHandler) Login(c *fiber.Ctx) error {

	// Parse request
	var req *requests.UserLoginRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	createPayload := entities.User{
		Email:    req.Email,
		Password: req.Password,
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

	return c.Status(fiber.StatusOK).JSON(responses.UserLoginResponse{
		ID:    user.UserId,
		Name:  user.Fname + " " + user.Lname,
		Email: user.Email,
		Token: token,
	})
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
	req := new(requests.UpdateUserRequest)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
			"log":     err.Error(),
		})
	}

	userId := utils.GetUserIDFromJWT(c)
	payload := entities.User{
		Email:       req.Email,
		Fname:       req.Fname,
		Lname:       req.Lname,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
	}

	user, err := uh.service.UpdateUser(c.Context(), userId, &payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successful update user",
		"payload": fiber.Map{
			"user": responses.UserDefaultResponse{
				ID:          user.UserId,
				Name:        user.Fname + " " + user.Lname,
				Email:       user.Email,
				PhoneNumber: user.PhoneNumber,
				Address:     user.Address,
			},
		},
	})
}
