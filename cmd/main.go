package main

import (
	"ps-go-fiber/blog/config"
	"ps-go-fiber/blog/internal/pages"
	"ps-go-fiber/blog/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/template/html/v3"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()

	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	engine := html.New("./html", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(slogfiber.New(customLogger))

	app.Use(recover.New())

	pages.NewHandler(app)

	app.Listen(":3000")
}
