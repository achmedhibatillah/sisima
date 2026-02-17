package testutil

import (
	"github.com/dvvnFrtn/sisima/internal/app"
	"github.com/gofiber/fiber/v3"
)

// SetupTestApp membuat instance Fiber App untuk unit test
func SetupTestApp() *fiber.App {
	return app.New(app.Config{
		EnableLogger: false, // matikan logger di test
	})
}
