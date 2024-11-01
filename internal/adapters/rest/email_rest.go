package rest

import (
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

type EmailHandler struct{
	usecase usecases.EmailUseCase
}

func ProvideEmailRestHandler(usecase usecases.EmailUseCase) *EmailHandler{
	return &EmailHandler{
		usecase: usecase,
	}
}

func (eh *EmailHandler) SendMailToAlertExpiredTransaction(c *fiber.Ctx) error{
	if err := eh.usecase.GetExpiredTransactionThisWeek(c.Context()); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}