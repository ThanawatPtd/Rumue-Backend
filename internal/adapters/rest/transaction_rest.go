package rest

import (
	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/requests"
	"github.com/ThanawatPtd/SAProject/domain/responses"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/gofiber/fiber/v2"
)

type TransactionRestHandler struct {
	service usecases.TransactionUseCase
}

func ProvideTransactionRestHandler(service usecases.TransactionUseCase) *TransactionRestHandler {
	return &TransactionRestHandler{
		service: service,
	}
}

func (th *TransactionRestHandler) CreateTransaction(c *fiber.Ctx) error {
	req := requests.CreateTransactionRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
			"log": err,
		})
	}

	createPayload := entities.Transaction{}
	if err := utils.MappingParser(&req, &createPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"log": err,
		})
	}

	userId := utils.GetUserIDFromJWT(c)
	vehicleId := c.Params("id")

	transaction, err := th.service.CreateTransaction(c.Context(), userId, vehicleId, &createPayload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"log": err,
		})
	}

	responseTransaction := responses.DefaultTransactionResponse{}
	if err = utils.MappingParser(transaction, &responseTransaction); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"log": err,
		})
	}

	
	return c.JSON(fiber.Map{
		"transaction": responseTransaction,
	})

}