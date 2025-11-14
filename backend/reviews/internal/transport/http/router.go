package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"reviews/internal/transport/http/handlers"
	"reviews/internal/transport/http/middleware"
)

func SetupRouters(handler handlers.Handler, md middleware.Middleware, env string) *fiber.App {
	app := fiber.New()
	app.Use(setupCORS(env), logger.New())

	app.Get("/:id", handler.GetByServiceID)
	app.Get("/count/:id", handler.GetCount)
	app.Get("/user/:id", handler.GetByUserID)
	app.Get("/rating/:id", handler.GetAverageRating)

	auth := app.Use(md.AuthMiddleware())
	{
		auth.Post("/", handler.CreateReview)
		auth.Put("/:id", handler.UpdateComment)
		auth.Delete("/:id", handler.DeleteComment)
	}

	return app
}

func setupCORS(env string) fiber.Handler {
	var allowOrigins []string
	switch env {
	case "prod":
		allowOrigins = append(allowOrigins, "http://94.241.170.57")
	case "local":
		allowOrigins = append(allowOrigins, "http://127.0.0.1:5173", "http://localhost:5173")

	}

	return cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length", "Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           86400,
	})
}
