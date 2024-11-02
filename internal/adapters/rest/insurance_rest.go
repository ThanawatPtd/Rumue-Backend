package rest

import (
	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/requests"
	"github.com/ThanawatPtd/SAProject/domain/responses"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/gofiber/fiber/v2"
)

type InsuranceHandler struct{
	usecase usecases.InsuranceUseCase
}

func ProvideInsuranceRestHandler(usecase usecases.InsuranceUseCase) *InsuranceHandler{
	return &InsuranceHandler{
		usecase: usecase,
	}
}

func (ih *InsuranceHandler) GetInsurance(c *fiber.Ctx) error{
	req := new(requests.GetInsuranceRequest)
	if err := c.BodyParser(req); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	payload := &entities.Insurance{
		Brand: req.Brand,
		Model: req.Model,
		Year: req.Year,
	}

	selected, err := ih.usecase.GetInsurance(c.Context(), payload)
	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Insurance not found",
		})
	}

	reponse := new(responses.GetInsuranceResponse)
	if err := utils.MappingParser(selected,reponse); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(reponse)
}


func (ih *InsuranceHandler) GetInsurances(c *fiber.Ctx) error{
	selected, err := ih.usecase.GetInsurances(c.Context())
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(selected)
}