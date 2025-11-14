package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	cat "services/intertal/transport/http/handlers/category"
	"services/intertal/transport/http/handlers/service"
	"services/intertal/transport/http/middleware"
)

func SetupRouters(service service.Handler, category cat.Handler, md middleware.Middleware, env string) *fiber.App {
	app := fiber.New()
	app.Use(setupCORS(env), logger.New())

	ser := app.Group("/services")
	{
		ser.Get("/", service.FindAll)
		ser.Get("/photo/:id", service.FindPhotosServices)
		ser.Get("/:id", service.FindByID)
	}

	cat := app.Group("/categories")
	{
		cat.Get("/", category.FindAll)
		cat.Get("/:id", category.FindByID)
	}

	// Для админа
	serAdmin := app.Group("/services").Use(md.RoleMiddleware())
	{
		serAdmin.Post("/", service.Create)
		serAdmin.Delete("/:id", service.Delete)
		serAdmin.Patch("/:id", service.Update)
	}

	catAdmin := app.Group("/categories").Use(md.RoleMiddleware())
	{
		catAdmin.Post("/", category.Create)
		catAdmin.Delete("/:id", category.Delete)
		catAdmin.Patch("/:id", category.Update)
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
