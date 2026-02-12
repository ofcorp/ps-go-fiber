package main

import (
	"ps-go-fiber/blog/config"
	"ps-go-fiber/blog/internal/pages"
	"ps-go-fiber/blog/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()

	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()

	app.Use(slogfiber.New(customLogger))

	app.Use(recover.New())

	pages.NewHandler(app)

	app.Listen(":3000")
}
