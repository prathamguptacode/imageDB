package src

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/prathamguptacode/imageDB/src/services"
)

func Server() {

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello world!! welcome to GO imageDB"})
	})

	app.Post("/file", services.Upload)
	app.Get("/file/:file", services.Retrieve)

	log.Fatal(app.Listen(":3000"))

}
