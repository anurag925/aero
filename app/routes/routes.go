package routes

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/anurag925/aero/app/handlers"
	v1 "github.com/anurag925/aero/app/handlers/api/v1"
	"github.com/anurag925/aero/config"
	"github.com/anurag925/aero/core"
	"github.com/labstack/echo/v4"
)

func Init() {
	server := core.Server()
	group := server.Group("")
	applicationController := handlers.NewApplicationHandler(group)
	v1ApplicationController := v1.NewV1ApplicationHandler(applicationController)
	v1UsersHandler := v1.NewUsersHandler(v1ApplicationController)
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
