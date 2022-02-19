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

type workspaceHandler struct {
	workspaceService service.WorkspaceService
}

func NewWorkspacesHandler(service service.WorkspaceService) WorkspacesHandler {
	return &workspaceHandler{workspaceService: service}
}

// GET /workspaces
func (h *workspaceHandler) Index(ctx echo.Context) error {
	ps, err := h.workspaceService.GetAll(ctx.Request().Context())
	if err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.JSON(http.StatusOK, ps)
}

// GET /workspaces/:id
func (h *workspaceHandler) Show(ctx echo.Context) error {
	id := ctx.Param("id")
	p, err := h.workspaceService.Get(ctx.Request().Context(), id)
	if err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.JSON(http.StatusOK, p)
}

// POST /workspaces
func (h *workspaceHandler) Create(ctx echo.Context) error {
	type Create struct {
		Db string `json:"db"`
	}

	var create Create
	if err := ctx.Bind(&create); err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	p, err := h.workspaceService.Create(ctx.Request().Context(), create.Db)
	if err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.JSON(http.StatusOK, p)
}

// DELETE /workspaces/:id
func (h *workspaceHandler) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := h.workspaceService.Delete(ctx.Request().Context(), id); err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// POST /workspaces/:id/query
func (h *workspaceHandler) Query(ctx echo.Context) error {
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

	r, err := h.workspaceService.Execute(ctx.Request().Context(), id, q.Query)
	if err != nil {
		ctx.Logger().Error(err)
		res := ErrorResponse{Error: err.Error()}
		return ctx.JSON(http.StatusInternalServerError, res)
	}

	return ctx.JSON(http.StatusOK, r)
}
