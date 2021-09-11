package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/koyashiro/postgres-playground/backend/handler"
	"github.com/koyashiro/postgres-playground/backend/repository"
	"github.com/koyashiro/postgres-playground/backend/service"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO: replace DI
	pr := repository.NewPlaygroundRepository()
	cr, err := repository.NewContainerRepository()
	if err != nil {
		panic(err)
	}
	ps := service.NewPlaygroundService(pr, cr)
	ph := handler.NewPlaygroundsHandler(ps)

	// Routes
	e.GET("/playgrounds", ph.GetPlaygrounds)
	e.GET("/playgrounds/:id", ph.GetPlayground)
	e.POST("/playgrounds", ph.PostPlayground)
	e.DELETE("/playgrounds/:id", ph.DeletePlayground)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
