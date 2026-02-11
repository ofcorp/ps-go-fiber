package main

import (
	"ps-go-fiber/blog/config"
	"ps-go-fiber/blog/internal/pages"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config.Init()
	// dbConf := config.NewDatabaseConfig()
	app := fiber.New()

	pages.NewHandler(app)

	app.Listen(":3000")
}
