package rest

import (
	"github.com/ThanawatPtd/SAProject/domain/entities"
	"github.com/ThanawatPtd/SAProject/domain/requests"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
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

	createPayload := entities.Vehicle(req)

	newVehicle, err := v.vehicleService.CreateVehicle(c.Context(), &createPayload)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Internal server error",
			"log":     err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Create Vehicle ",
		"payload": fiber.Map{
			"vehicle": newVehicle,
		},
	})
}
