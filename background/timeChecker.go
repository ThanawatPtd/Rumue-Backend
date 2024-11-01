package background

import (
	"log"
	"time"

	"github.com/ThanawatPtd/SAProject/internal/adapters/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func DailyMonitor(app *fiber.App, handler *rest.Handler){
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	for{
        time.Sleep(24*time.Hour)
		if err := handler.Email.SendMailToAlertExpiredTransaction(c); err !=nil{
			log.Fatal("Email service error")
		}
	}
}