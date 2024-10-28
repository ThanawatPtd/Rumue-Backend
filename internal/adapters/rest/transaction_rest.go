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
			"log":     err,
		})
	}
	createPayload := entities.Transaction{}
	if err := utils.MappingParser(&req, &createPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"log": err,
		})
	}

	jwt := utils.GetJWTFromContext(c)
	vehicleId := c.Params("id")
	transaction, err := th.service.CreateTransaction(c.Context(), jwt.UserID, vehicleId, &createPayload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"log": err,
		})
	}

	responseTransaction := responses.CreateTransactionResponse{}
	if err = utils.MappingParser(transaction, &responseTransaction); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"log": err,
		})
	}

	return c.JSON(fiber.Map{
		"transaction": responseTransaction,
	})

}

func (th *TransactionRestHandler) CheckHistory(c *fiber.Ctx) error {
	jwt := utils.GetJWTFromContext(c)
	userVehicleTransactions, err := th.service.CheckHistory(c.Context(), jwt.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"log": err,
		})
	}

	var responseTransaction []responses.DefaultTransactionResponse
	for _, value := range userVehicleTransactions {
		var transaction responses.DefaultTransactionResponse
		if err = utils.MappingParser(&value, &transaction); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"log": err,
			})
		}
		responseTransaction = append(responseTransaction, transaction)
	}
	return c.JSON(fiber.Map{
		"transactions": responseTransaction,
	})
}
func (th *TransactionRestHandler) FindInsuranceToday(c *fiber.Ctx) error {
	response, err := th.service.FindTodayInsurances(c.Context())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success Get Transaction",
		"payload": fiber.Map{
			"Transactions": response,
		},
	})
}

func (th *TransactionRestHandler) UpdateTransaction(c *fiber.Ctx) error {
	req := requests.UpdateTransactionRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
			"log":     err,
		})
	}
	createPayload := entities.Transaction{}
	if err := utils.MappingParser(&req, &createPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"log": err,
		})
	}
	err := th.service.UpdateTransaction(c.Context(), &createPayload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Update transaction successful",
	})
}

func (th *TransactionRestHandler) GetUserVehicleTransactionByID(c *fiber.Ctx) error {
	transactionID := c.Params("id")
	userVehicleTransaction, err := th.service.FindTransactionByID(c.Context(), transactionID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	var response responses.DefaultTransactionResponse
	if err := utils.MappingParser(userVehicleTransaction, &response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"log": err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successful get user",
		"payload": fiber.Map{
			"transaction": response,
		},
	})
}
func (th *TransactionRestHandler) SumThreeMonthIncome(c *fiber.Ctx) error {

	income, err := th.service.SumThreeMonthIncome(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	var response responses.IncomeResponse

	if err := utils.MappingParser(income, &response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"log": err,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success",
		"Income":  income,
	})
}
