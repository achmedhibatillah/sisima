package app

import (
	"github.com/dvvnFrtn/sisima/internal/logger"
	route "github.com/dvvnFrtn/sisima/internal/routes"
	"github.com/gofiber/fiber/v3"
)

type Config struct {
	EnableLogger bool
}

func New(cfg Config) *fiber.App {
	app := fiber.New()

	if cfg.EnableLogger {
		app.Use(logger.HTTPLogger())
	}

	route.RegisterRoutes(app)

	return app
}
