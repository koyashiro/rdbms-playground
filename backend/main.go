package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/koyashiro/postgres-playground/backend/handler"
	"github.com/koyashiro/postgres-playground/backend/repository"
	"github.com/koyashiro/postgres-playground/backend/service"
)

const port = "1323"

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// TODO: replace DI
	cr, err := repository.NewContainerRepository()
	if err != nil {
		panic(err)
	}
	rr := repository.NewRDBMSRepository()
	ps := service.NewWorkspaceService(cr, rr)
	ph := handler.NewWorkspacesHandler(ps)

	// Routes
	e.GET("/workspaces", ph.GetWorkspaces)
	e.GET("/workspaces/:id", ph.GetWorkspace)
	e.POST("/workspaces", ph.PostWorkspace)
	e.DELETE("/workspaces/:id", ph.DeleteWorkspace)
	e.POST("/workspaces/:id/query", ph.ExecuteQuery)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
