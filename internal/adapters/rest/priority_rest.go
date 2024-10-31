package rest

import (
	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/responses"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/gofiber/fiber/v2"
)


type PriorityHandler struct{
	usecase usecases.PriorityUseCase
}

func ProvidePriorityRestHandler(usecase usecases.PriorityUseCase) *PriorityHandler{
	return &PriorityHandler{
		usecase: usecase,
	}
}

func (ph *PriorityHandler) InitializePriority(c *fiber.Ctx) error{
	initRange := []string{"1-3000", "4000-7000", "7000-i"}
	initRate := []float64{0.8, 1.0, 1.2}


	for inx, val := range initRange {
		if _, err := ph.usecase.CreatePriorityByID(c.Context(), &entities.Priority{
			ID: val,
			Rate: initRate[inx],
		}); err != nil{
			defer panic("Initialize Error")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error" : "Initialize Error",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "created successful",
	})

}


func (ph *PriorityHandler) GetPriorityRateByID(c *fiber.Ctx) error{
	id := c.Params("id")

	selected, err := ph.usecase.GetPriorityByID(c.Context(), id)

	if err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err,
		})
	}

	temp := new(responses.PriorityResponse)
	if err := utils.MappingParser(selected, temp); err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":err,
		})
	}

	return c.JSON(temp)
}