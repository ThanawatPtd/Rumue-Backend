package router

import (
	"fmt"

	"github.com/ThanawatPtd/SAProject/internal/adapters/rest"
	"github.com/gofiber/fiber/v2"
)

func RegisterApiRouter(app *fiber.App, handler *rest.Handler) {
	api := app.Group("/")

	user := api.Group("/user")

	fmt.Print(api)
	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
		})
	})

	user.Get("/:id",handler.User.GetUserByID)

	user.Post("/create", handler.User.CreateUser)
	// user.Update("/update/:id", )

}
