package rest

import (
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/gofiber/fiber/v2"
)

type UserRestHandler struct {
	userUseCase usecases.UserUseCase
}


func NewUserRestHandler(userUseCase usecases.UserUseCase) *UserRestHandler{
	return &UserRestHandler{userUseCase: userUseCase}
}


func (uh *UserRestHandler) CreateUser(c *fiber.Ctx) error{
	user := new(dbmodel.User)


	if err:=c.BodyParser(user);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Invalid data",
		})
	}

	selectedUser, err :=uh.userUseCase.CreateUser(user)

	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message":err,
			})
	}


	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":selectedUser,
	})
}