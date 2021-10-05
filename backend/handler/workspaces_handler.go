package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/koyashiro/rdbms-playground/backend/service"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type WorkspacesHandler interface {
	// GET /workspaces
	Index(c echo.Context) error

	// GET /workspaces/:id
	Show(c echo.Context) error

	// POST /workspaces
	Create(c echo.Context) error

	// DELETE /workspaces/:id
	Delete(c echo.Context) error

	// POST /workspaces/:id/query
	Query(c echo.Context) error
}

type WorkspacesHandlerImpl struct {
	workspaceService service.WorkspaceService
}

func NewWorkspacesHandler(service service.WorkspaceService) WorkspacesHandler {
	return &WorkspacesHandlerImpl{workspaceService: service}
}

// GET /workspaces
func (h *WorkspacesHandlerImpl) Index(ctx echo.Context) error {
	ps, err := h.workspaceService.GetAll()
	if err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.JSON(http.StatusOK, ps)
}

// GET /workspaces/:id
func (h *WorkspacesHandlerImpl) Show(ctx echo.Context) error {
	id := ctx.Param("id")
	p, err := h.workspaceService.Get(id)
	if err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.JSON(http.StatusOK, p)
}

// POST /workspaces
func (h *WorkspacesHandlerImpl) Create(ctx echo.Context) error {
	type Create struct {
		Db string `json:"db"`
	}

	var create Create
	if err := ctx.Bind(&create); err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	p, err := h.workspaceService.Create(create.Db)
	if err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.JSON(http.StatusOK, p)
}

// DELETE /workspaces/:id
func (h *WorkspacesHandlerImpl) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := h.workspaceService.Delete(id); err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// POST /workspaces/:id/query
func (h *WorkspacesHandlerImpl) Query(ctx echo.Context) error {
	id := ctx.Param("id")

	type Query struct {
		Query string `json:"query"`
	}

	var q Query
	if err := ctx.Bind(&q); err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	r, err := h.workspaceService.Execute(id, q.Query)
	if err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.JSON(http.StatusOK, r)
}
