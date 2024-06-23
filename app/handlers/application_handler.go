package handlers

import (
	"context"

	"github.com/labstack/echo/v4"
)

type ApplicationHandler struct {
	group *echo.Group
}

func NewApplicationHandler(group *echo.Group) *ApplicationHandler {
	return &ApplicationHandler{group: group}
}

func (h *ApplicationHandler) Group() *echo.Group {
	return h.group
}

func (h *ApplicationHandler) Ctx(c echo.Context) context.Context {
	return c.Request().Context()
}
