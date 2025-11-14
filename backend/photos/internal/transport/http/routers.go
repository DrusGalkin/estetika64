package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
	"photos/internal/transport/http/handlers"
)

func SetupRouters(handler handlers.Handler, env string) *fiber.App {
	app := fiber.New()
	app.Use(setupCORS(env), logger.New())

	app.Get("/*", static.New(handlers.UPLOADS_PATH))
	app.Get("/photos/:id", handler.GetByServiceID)
	app.Get("/gallery", handler.AllGallery)
	app.Get("/gallery/two", handler.TwoPhotosGallery)

	app.Post("/uploads/:serId/:index", handler.Upload)
	app.Patch("/edit-index/:serId/:index", handler.EditIndex)

	app.Post("/teammates/uploads", handler.UploadTeammatePhoto)

	app.Post("/gallery", handler.UploadGallery)
	app.Delete("/gallery/:id", handler.DeleteGallery)

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
