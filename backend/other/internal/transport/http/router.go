package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"other/internal/transport/http/handlers"
	"other/internal/transport/http/middleware"
)

func SetupRouter(handler *handlers.OtherHandler, middle middleware.Middleware, env string) *fiber.App {
	app := fiber.New()
	app.Use(setupCORS(env), logger.New())

	about := app.Group("/about")
	{
		about.Get("/", handler.GetAbout)
		admin := about.Use(middle.RoleMiddleware())
		{
			admin.Post("/", handler.CreateAbout)
			admin.Put("/:id", handler.UpdateAbout)
			admin.Delete("/:id", handler.DeleteAbout)
		}

	}

	contacts := app.Group("/contacts")
	{
		contacts.Get("/", handler.GetAllContacts)
		admin := contacts.Use(middle.RoleMiddleware())
		{
			admin.Post("/", handler.CreateContact)
			admin.Put("/", handler.UpdateContact)
			admin.Delete("/:id", handler.DeleteContact)
		}
	}

	teammates := app.Group("/teammates")
	{
		teammates.Get("/", handler.GetAllTeammates)
		admin := teammates.Use(middle.RoleMiddleware())
		{
			admin.Post("/", handler.CreateTeammate)
			admin.Put("/:id", handler.UpdateTeammate)
			admin.Delete("/:id", handler.DeleteTeammate)
		}
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
