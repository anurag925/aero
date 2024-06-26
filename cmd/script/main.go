package main

import (
	"log/slog"

	"github.com/anurag925/aero/cmd"
	"github.com/anurag925/aero/config"
	"github.com/anurag925/aero/core"
)

func main() {
	slog.Info("Starting the application from the main...")
	flags := cmd.ParseFlags()
	slog.Info("Application is starting on", slog.Any("environment", config.Env(config.Environment(*flags.Env))), slog.Int("port", *flags.Port))
	core.Script()

}
