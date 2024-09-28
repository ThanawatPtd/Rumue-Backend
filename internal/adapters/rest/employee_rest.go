package rest

import (
	"github.com/ThanawatPtd/SAProject/domain/requests"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db/dbmodel"
	"github.com/emicklei/pgtalk/convert"
	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct{
	employee usecases.EmployeeUseCase
}

func ProvideEmployeeRestHandler(employee usecases.EmployeeUseCase) *EmployeeHandler{
	return &EmployeeHandler{
		employee: employee,
	}
}

func (eh *EmployeeHandler) GetEmployeeByID(c *fiber.Ctx) error{

	id := convert.StringToUUID(c.Params("id"))

	response, err := eh.employee.GetByID(c.Context(), &id)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"log":err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"Successful get employee",
		"payload":fiber.Map{
			"employee": response,
		},
	})
}


func (eh *EmployeeHandler) GetEmployees(c *fiber.Ctx) error{
	response, err := eh.employee.ListAll(c.Context())

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"log":err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":"Successful get employees",
		"payload":fiber.Map{
			"employees":response,
		},
	})
}


func (eh *EmployeeHandler) CreateEmployee(c *fiber.Ctx) error{
	rq := new(requests.CreateEmployeeRequest)

	if err := c.BodyParser(rq); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
			"log":err.Error(),
		})
	}


	payload := dbmodel.CreateEmployeeParams(*rq)

	response, err := eh.employee.Save(c.Context(), &payload)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"log":err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"Successful created employee",
		"payload":fiber.Map{
			"employee": response,
		},
	})
}


func (eh *EmployeeHandler) DeleteByID(c *fiber.Ctx) error{
	id := convert.StringToUUID(c.Params("id"))


	if err := eh.employee.Delete(c.Context(), &id); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
			"log":err.Error(),
		})
	}


	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message":"Successful delete employee",
	})
}