package main

import (
	"os"

	"github.com/dvvnFrtn/sisima/internal/app"
	"github.com/dvvnFrtn/sisima/internal/config"
)

func main() {
	config.ConnectDatabase()

	app := app.New(app.Config{
		EnableLogger: true,
	})

	app.Listen(":" + os.Getenv("SERVER_PORT"))
}
