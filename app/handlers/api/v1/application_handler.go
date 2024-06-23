package v1

import (
	"github.com/anurag925/aero/app/handlers"
	"github.com/anurag925/aero/app/views"
	"github.com/labstack/echo/v4"
)

type v1ApplicationHandler struct {
	*handlers.ApplicationHandler
	group *echo.Group
}

func NewV1ApplicationHandler(app *handlers.ApplicationHandler) *v1ApplicationHandler {
	return &v1ApplicationHandler{app, app.Group()}
}

func (h *v1ApplicationHandler) Index(c echo.Context) error {
	return c.Render(200, "hello_word", views.Hello("word"))
}

func (h *v1ApplicationHandler) Group() *echo.Group {
	return h.group.Group("api/v1")
}
