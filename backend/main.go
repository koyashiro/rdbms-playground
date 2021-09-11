package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/koyashiro/postgres-playground/backend/handler"
	"github.com/koyashiro/postgres-playground/backend/repositories"
	"github.com/koyashiro/postgres-playground/backend/services"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO: replace DI
	playgroundRepository := repositories.NewPlaygroundRepository()
	playgroundServices := services.NewPlaygroundService(&playgroundRepository)
	playgroundsHandler := handler.NewPlaygroundsHandler(playgroundServices)

	// Routes
	e.GET("/playgrounds", playgroundsHandler.GetPlaygrounds)
	e.GET("/playgrounds/:id", playgroundsHandler.GetPlayground)
	e.POST("/playgrounds", playgroundsHandler.PostPlayground)
	e.DELETE("/playgrounds/:id", playgroundsHandler.DeletePlayground)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
