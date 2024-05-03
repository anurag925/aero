package v1

import "github.com/labstack/echo/v4"

type usersHandler struct {
	*v1ApplicationHandler
}

func NewUsersHandler(app *v1ApplicationHandler) *usersHandler {
	return &usersHandler{app}
}

func (h *usersHandler) login(c echo.Context) error {
	return nil
}

func (h *usersHandler) register(c echo.Context) error {
	return nil
}

func (h *usersHandler) Routes() {
	group := h.Group().Group("/users")
	group.POST("/login", h.login)
	group.POST("/register", h.register)
}
