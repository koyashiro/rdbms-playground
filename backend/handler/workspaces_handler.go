package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/koyashiro/postgres-playground/backend/service"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type WorkspacesHandler interface {
	GetWorkspaces(c echo.Context) error
	GetWorkspace(c echo.Context) error
	PostWorkspace(c echo.Context) error
	DeleteWorkspace(c echo.Context) error
	ExecuteQuery(c echo.Context) error
}

type WorkspacesHandlerImpl struct {
	workspaceService service.WorkspaceService
}

func NewWorkspacesHandler(service service.WorkspaceService) WorkspacesHandler {
	return &WorkspacesHandlerImpl{workspaceService: service}
}

func (h *WorkspacesHandlerImpl) GetWorkspaces(c echo.Context) error {
	ps, err := h.workspaceService.GetAll()
	if err != nil {
		c.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, ps)
}

func (h *WorkspacesHandlerImpl) GetWorkspace(c echo.Context) error {
	id := c.Param("id")
	p, err := h.workspaceService.Get(id)
	if err != nil {
		c.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, p)
}

func (h *WorkspacesHandlerImpl) PostWorkspace(c echo.Context) error {
	type Create struct {
		Db string `json:"db"`
	}

	var create Create
	if err := c.Bind(&create); err != nil {
		c.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	p, err := h.workspaceService.Create(create.Db)
	if err != nil {
		c.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, p)
}

func (h *WorkspacesHandlerImpl) DeleteWorkspace(c echo.Context) error {
	id := c.Param("id")
	if err := h.workspaceService.Destroy(id); err != nil {
		c.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *WorkspacesHandlerImpl) ExecuteQuery(c echo.Context) error {
	id := c.Param("id")

	type Query struct {
		Query string `json:"query"`
	}

	var q Query
	if err := c.Bind(&q); err != nil {
		c.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	r, err := h.workspaceService.Execute(id, q.Query)
	if err != nil {
		c.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, r)
}
