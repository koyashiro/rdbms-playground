package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/playgrounds", postPlayground)
	e.DELETE("/playgrounds/:id", deletePlayground)
	e.POST("/playgrounds/:id/execute", executeQuery)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type PlaygroundCreationResponse struct {
	Id string `json:"id"`
}

func postPlayground(c echo.Context) error {
	id := uuid.NewString()
	c.Logger().Info("create playground: " + id)
	res := PlaygroundCreationResponse{
		Id: uuid.NewString(),
	}
	return c.JSON(http.StatusOK, res)
}

func deletePlayground(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Info("delete playground: " + id)
	return c.JSON(http.StatusNoContent, nil)
}

type ExecuteQueryRequest struct {
	Id    string `json:"id"`
	Query string `json:"query"`
}

type ExecuteQueryResponse struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func executeQuery(c echo.Context) error {
	id := c.Param("id")
	req := new(ExecuteQueryRequest)
	if err := c.Bind(req); err != nil {
		c.Logger().Error("invalid parameter")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid parameter"})
	}

	res := ExecuteQueryResponse{
		Result: "execution result",
	}
	c.Logger().Info("execute query: " + id)
	return c.JSON(http.StatusOK, res)
}
