package rest

import (
	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/requests"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	employee usecases.EmployeeUseCase
}

func ProvideEmployeeRestHandler(employee usecases.EmployeeUseCase) *EmployeeHandler {
	return &EmployeeHandler{
		employee: employee,
	}
}

func (eh *EmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	rq := new(requests.CreateEmployeeRequest)

	if err := c.BodyParser(rq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"log":     err.Error(),
		})
	}

	payload := entities.Employee{}
	if err := utils.MappingParser(rq, &payload); err != nil {
		return err
	}

	response, err := eh.employee.Save(c.Context(), &payload)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successful created employee",
		"payload": fiber.Map{
			"employee": response,
		},
	})
}
