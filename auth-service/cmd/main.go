package main

import (
	"github.com/saufiroja/sosmed-app/auth-service/internal/app"
)

func main() {
	app := app.NewApp()
	app.RunApp()
}
