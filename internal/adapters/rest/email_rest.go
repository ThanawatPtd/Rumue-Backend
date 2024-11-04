package rest

import (
	"encoding/base64"
	"io"
	"log"

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

func (eh *EmailHandler) SendMailToAlertReceipt(c *fiber.Ctx) error{
	transactionID := c.FormValue("transactionID")

	if transactionID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing transaction ID")
	}

	file, err := c.FormFile("receiptFile")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Could not get uploaded file")
	}

	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not open uploaded file")
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not read file content")
	}

	encodedFile := base64.StdEncoding.EncodeToString(fileBytes)

	err = eh.usecase.SendReceipt(c.Context(),encodedFile, file.Filename, transactionID)

	if err != nil {
		log.Println("Failed to send email:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Could not send email")
	}

	return c.SendStatus(fiber.StatusNoContent)
}