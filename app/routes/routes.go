package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/anurag925/aero/app/handlers"
	v1 "github.com/anurag925/aero/app/handlers/api/v1"
	"github.com/anurag925/aero/app/repositories/mysql"
	"github.com/anurag925/aero/app/services/impl"
	"github.com/anurag925/aero/config"
	"github.com/anurag925/aero/core"
	"github.com/labstack/echo/v4"
)

func Init() {
	server := core.Server()
	group := server.Group("/")

	// health check api
	group.GET("up", func(c echo.Context) error { return c.JSON(http.StatusOK, "up") })

	applicationController := handlers.NewApplicationHandler(group)

	// all controllers
	v1ApplicationController := v1.NewV1ApplicationHandler(*applicationController)

	// this is an example of dependency injection
	v1UsersHandler := v1.NewUsersHandler(*v1ApplicationController, impl.NewUsersService(mysql.NewUsersRepository(core.DB())))
	v1UsersHandler.Routes()
	printRoutes(server)
}

// printRoutes writes the routes to a file for debugging
func printRoutes(s *echo.Echo) error {
	if s == nil {
		return errors.New("no echo server given")
	}

	if !config.IsDevelopment() {
		return nil
	}

	data, err := json.MarshalIndent(s.Routes(), "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile("routes.json", data, 0644); err != nil {
		return err
	}
	return nil
}
