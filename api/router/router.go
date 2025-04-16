package router

import(
	"github.com/gofiber/fiber/v2"
	"ai-pdf-chat/api/handler"
	"ai-pdf-chat/middleware"
)

func AppRoute(app *fiber.App) {

	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group("/api")
	api.Post("/login", handler.LoginUser)

	file := app.Group("/api/file", middleware.JWTAuthentication)
	file.Get("/list", handler.ListFiles)
	file.Post("/upload", handler.UploadFile)
	file.Post("/processed", handler.ProcessFile)
}