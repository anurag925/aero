package v1

import (
	"net/http"

	"github.com/anurag925/aero/app/models"
	"github.com/anurag925/aero/app/repositories/mongo"
	"github.com/anurag925/aero/app/services"
	"github.com/anurag925/aero/config"
	"github.com/anurag925/aero/core"
	"github.com/anurag925/aero/core/logger"
	"github.com/labstack/echo/v4"
)

type usersHandler struct {
	v1ApplicationHandler

	usersService services.UsersService
}

func NewUsersHandler(app v1ApplicationHandler, usersService services.UsersService) *usersHandler {
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

	rl := models.RequestLog{RequestID: "sadvx"}
	err := mongo.NewRequestLogRepositoryImpl(core.MongoDB(), config.Secrets().MongoDBName).Log(h.Ctx(c), &rl)
	if err != nil {
		logger.Error(h.Ctx(c), "Error logging request", "data", err)
		return err
	}
	logger.Info(h.Ctx(c), "Request logged", rl)
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
