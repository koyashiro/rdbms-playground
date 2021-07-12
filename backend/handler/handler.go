package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PlaygroundCreationResponse struct {
	Id string `json:"id"`
}

type Playground struct {
	Id string `json:"id"`
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

func PostPlayground(c echo.Context) error {
	id := uuid.NewString()
	c.Logger().Info("create playground: " + id)
	res := PlaygroundCreationResponse{
		Id: uuid.NewString(),
	}
	return c.JSON(http.StatusOK, res)
}

func GetPlayground(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Info("show playground: " + id)
	playground := Playground{
		Id: id,
	}
	return c.JSON(http.StatusOK, playground)
}

func DeletePlayground(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Info("delete playground: " + id)
	return c.JSON(http.StatusNoContent, nil)
}

func ExecuteQuery(c echo.Context) error {
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
