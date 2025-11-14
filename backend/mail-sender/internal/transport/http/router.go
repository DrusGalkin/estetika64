package http

import (
	"github.com/DrusGalkin/go-mail-sender/internal/transport/http/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouters(handler handlers.Handler, env string) *fiber.App {
	app := fiber.New()
	app.Use(setupCORS(env), logger.New())

	app.Post("/send", handler.SendToEmail)
	app.Post("/confirm", handler.ConfirmEmail)
	app.Post("/statement", handler.SendStatement)

	return app
}

func setupCORS(env string) fiber.Handler {
	var allowOrigins string
	switch env {
	case "prod":
		allowOrigins = "http://estetica64.ru,http://www.estetica64.ru"
	case "local":
		allowOrigins = "http://127.0.0.1:5173,http://localhost:5173"
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		ExposeHeaders:    "Content-Length, Set-Cookie",
		AllowCredentials: true,
		MaxAge:           86400,
	})
}
