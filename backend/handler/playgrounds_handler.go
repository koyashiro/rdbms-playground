package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/koyashiro/postgres-playground/backend/service"
)

type ErrorResponse struct {
	Error string `json:"message"`
}

type PlaygroundsHandler interface {
	GetPlaygrounds(c echo.Context) error
	GetPlayground(c echo.Context) error
	PostPlayground(c echo.Context) error
	DeletePlayground(c echo.Context) error
}

type PlaygroundsHandlerImpl struct {
	playgroundService service.PlaygroundService
}

func NewPlaygroundsHandler(service service.PlaygroundService) PlaygroundsHandler {
	return &PlaygroundsHandlerImpl{playgroundService: service}
}

func (h *PlaygroundsHandlerImpl) GetPlaygrounds(c echo.Context) error {
	ps, err := h.playgroundService.GetAll()
	if err != nil {
		c.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, ps)
}

func (h *PlaygroundsHandlerImpl) GetPlayground(c echo.Context) error {
	id := c.Param("id")
	p, err := h.playgroundService.Get(id)
	if err != nil {
		c.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, p)
}

func (h *PlaygroundsHandlerImpl) PostPlayground(c echo.Context) error {
	p, err := h.playgroundService.Create("postgres")
	if err != nil {
		c.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, p)
}

func (h *PlaygroundsHandlerImpl) DeletePlayground(c echo.Context) error {
	id := c.Param("id")
	if err := h.playgroundService.Destroy(id); err != nil {
		c.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.NoContent(http.StatusNoContent)
}
