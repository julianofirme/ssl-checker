package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jfirme-sys/ssl-checker/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	router.SetupRouter(app)
	log.Fatal(app.Listen(":3000"))
}
