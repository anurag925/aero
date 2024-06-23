package v1

import (
	"net/http"

	"github.com/anurag925/aero/app/services"
	"github.com/labstack/echo/v4"
)

type usersHandler struct {
	*v1ApplicationHandler

	usersService services.UsersService
}

func NewUsersHandler(app *v1ApplicationHandler, usersService services.UsersService) *usersHandler {
	return &usersHandler{app, usersService}
}

type showParams struct {
	ID int64 `param:"id"`
}

func (h *usersHandler) show(c echo.Context) error {
	params := showParams{}
	if err := c.Bind(&params); err != nil {
		return err
	}
	user, err := h.usersService.FindUserById(h.Ctx(c), params.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *usersHandler) login(c echo.Context) error {
	return nil
}

func (h *usersHandler) register(c echo.Context) error {
	return nil
}

func (h *usersHandler) Routes() {
	group := h.Group().Group("/users")
	group.GET("/:id", h.show)
	group.POST("/login", h.login)
	group.POST("/register", h.register)
}
