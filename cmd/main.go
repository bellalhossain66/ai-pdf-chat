package main

import(
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"ai-pdf-chat/config"
	"ai-pdf-chat/api/router"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	app := fiber.New()
	router.AppRoute(app)

	log.Printf("==> 🌎  Listening on port %s. Visit http://localhost:%s/ in your browser.", os.Getenv("APP_PORT"), os.Getenv("APP_PORT"))
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Fail to start server %v", err)
	}
}