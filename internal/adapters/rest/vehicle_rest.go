package rest

import (
	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/requests"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/utils"
	"github.com/gofiber/fiber/v2"
)

type VehicleRestHandler struct {
	vehicleService usecases.VehicleUseCase
}

func ProvideVehicleHandler(vehicleUseCase usecases.VehicleUseCase) *VehicleRestHandler {
	return &VehicleRestHandler{
		vehicleService: vehicleUseCase,
	}
}

func (v *VehicleRestHandler) CreateVehicle(c *fiber.Ctx) error {

	var req requests.CreateVehicleRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"log":     err.Error(),
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	createPayload := &entities.Vehicle{}

	if err := utils.MappingParser(&req, createPayload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	jwt := utils.GetJWTFromContext(c)
	newVehicle, err := v.vehicleService.CreateVehicle(c.Context(), jwt.UserID, createPayload)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Create Vehicle Success",
		"payload": fiber.Map{
			"vehicleID": newVehicle.ID,
		},
	})
}

func (v *VehicleRestHandler) FindTemplate(c *fiber.Ctx) error {

	jwt := utils.GetJWTFromContext(c)
	vehicles, err := v.vehicleService.FindTemplate(c.Context(), jwt.UserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message":  "Find Success",
		"vehicles": vehicles,
	})
}
