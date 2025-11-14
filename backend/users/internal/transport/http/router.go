package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"users/internal/transport/http/handlers"
	"users/internal/transport/http/middleware"
)

func Setup(h handlers.Handler, md middleware.App, env string) *fiber.App {
	app := fiber.New()
	app.Use(setupCORS(env), logger.New())

	app.Post("/register", h.Register)
	app.Post("/refresh", h.Refresh)
	app.Post("/login", h.Login)
	app.Post("/valid", h.ValidToken)

	app.Get("/", h.GetAllUsers)
	app.Get("/:id", h.FindUser)
	app.Post("/check/:id", h.IsAdmin)
	app.Post("/email", h.FindUserByEmail)

	auth := app.Use(md.Auth())
	{
		auth.Patch("/password", h.UpdatePassword)
		auth.Patch("/email", h.UpdateEmail)
		auth.Delete("/", h.Delete)
		auth.Post("/out", h.Logout)
	}

	return app
}

func setupCORS(env string) fiber.Handler {
	var allowOrigins []string
	switch env {
	case "prod":
		allowOrigins = append(allowOrigins, "http://estetica64.ru", "http://www.estetica64.ru")
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
