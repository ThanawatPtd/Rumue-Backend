package rest

import (
	// "fmt"
	"log"

	"github.com/ThanawatPtd/SAProject/domain/requests"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/emicklei/pgtalk/convert"
	"github.com/gofiber/fiber/v2"
)

type UserRestHandler struct {
	userUseCase usecases.UserUseCase
}

func ProvideUserRestHandler(userUseCase usecases.UserUseCase) *UserRestHandler {
	return &UserRestHandler{userUseCase: userUseCase}
}

func (uh *UserRestHandler) GetUsers(c *fiber.Ctx) error {
	list, err := uh.userUseCase.GetUsers(c.Context())

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
	id := convert.StringToUUID(c.Params("id"))

	user, err := uh.userUseCase.GetUserByID(c.Context(), &id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"log":     err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successful get user",
		"payload": fiber.Map{
			"user": user,
		},
	})
}

func (uh *UserRestHandler) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
			"log":     "empty email in subpath",
		})
	}

	response, err := uh.userUseCase.GetByEmail(c.Context(), &email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successful get user",
		"payload": fiber.Map{
			"user": response,
		},
	})
}

func (uh *UserRestHandler) CreateUser(c *fiber.Ctx) error {
	var req requests.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		log.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	createPayload := dbmodel.CreateUserParams(req)
	selectedUser, err := uh.userUseCase.CreateUser(c.Context(), &createPayload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err,
			})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": selectedUser,
	})
}

func (uh *UserRestHandler) DeleteByID(c *fiber.Ctx) error {
	id := convert.StringToUUID(c.Params("id"))

	if err := uh.userUseCase.DeleteByID(c.Context(), &id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)

}

// func (uh *UserRestHandler) UpdateUser(c *fiber.Ctx) error{
// 	rq := new(requests.UpdateUserRequest)

// 	if err:= c.BodyParser(rq); err != nil{
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message":"Bad request",
// 			"log":err.Error(),
// 		})
// 	}

// 	payload := dbmodel.UpdateUserParams(*rq)

// 	id := convert.StringToUUID(c.Params("id"))

// 	user, err := uh.userUseCase.UpdateUser(c.Context(), &id, &payload)

// 	if err != nil{
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message":"Internal server error",
// 			"log":err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message":"Successful update user",
// 		"payload":fiber.Map{
// 			"user":user,
// 		},
// 	})
// }
