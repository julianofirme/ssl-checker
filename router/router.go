package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jfirme-sys/ssl-checker/handler"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Get("/hello", handler.Hello)
	api.Post("/ssl-check", handler.SSLCheck)
}
