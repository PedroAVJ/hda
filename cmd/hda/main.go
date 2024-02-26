package main

import (
	"github.com/PedroAVJ/hda/pkg/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.GET("/", handlers.HomeHandler)
	app.Logger.Fatal(app.Start(":4000"))
}
