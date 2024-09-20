package rest

<<<<<<< HEAD
import (
	"log"

	"github.com/ThanawatPtd/SAProject/domain/requests"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/gofiber/fiber/v2"
)

type UserRestHandler struct {
	userUseCase usecases.UserUseCase
=======
import "github.com/ThanawatPtd/SAProject/domain/usecases"

type UserRestHandler struct {
	userService usecases.UserUseCase
}

func ProvideUserRestHandler(userService usecases.UserUseCase) *UserRestHandler {
	return &UserRestHandler{
		userService: userService,
	}
>>>>>>> 6a08fd7 (Got stuck here need to wait for other complete)
}

func NewUserRestHandler(userUseCase usecases.UserUseCase) *UserRestHandler {
	return &UserRestHandler{userUseCase: userUseCase}
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

