package main

import (
	"github.com/koyashiro/postgres-playground/backend/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/playgrounds", handler.PostPlayground)
	e.GET("/playgrounds/:id", handler.GetPlayground)
	e.DELETE("/playgrounds/:id", handler.DeletePlayground)
	e.POST("/playgrounds/:id/execute", handler.ExecuteQuery)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
