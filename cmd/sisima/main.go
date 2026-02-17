package main

import (
	"os"

	"github.com/dvvnFrtn/sisima/internal/app"
	"github.com/dvvnFrtn/sisima/internal/config"
	model "github.com/dvvnFrtn/sisima/internal/models"
)

func main() {
	config.ConnectDatabase()
	model.Migrate()

	app := app.New(app.Config{
		EnableLogger: true,
	})

	app.Listen(":" + os.Getenv("SERVER_PORT"))
}
