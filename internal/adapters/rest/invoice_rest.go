package rest

import (
	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/requests"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

type InvoiceRestHandler struct {
	invoiceUseCase usecases.InvoiceUseCase
}

func ProvideInvoiceRestHandler(invoice usecases.InvoiceUseCase) *InvoiceRestHandler {
	return &InvoiceRestHandler{
		invoiceUseCase: invoice,
	}
}
func (ih *InvoiceRestHandler) GetInvoices(c *fiber.Ctx) error {
	list, err := ih.invoiceUseCase.GetInvoices(c.Context())

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Invoice not found",
			"log":     err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successful get invoice",
		"payload": fiber.Map{
			"invoice": list,
		},
	})
}

func (ih *InvoiceRestHandler) CreateInvoice(c *fiber.Ctx) error {
	var rq requests.CreateInvoiceRequest

	if err := c.BodyParser(&rq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"log":     err.Error(),
		})
	}

	payload := entities.Invoice(rq)

	response, err := ih.invoiceUseCase.CreateInvoice(c.Context(), &payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successful created invoice",
		"payload": fiber.Map{
			"invoice": response,
		},
	})
}
