package core

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/anurag925/aero/config"
	"github.com/anurag925/aero/core/initializers"
	"github.com/anurag925/aero/lib/utils/asynqwrapper"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type application struct {
	logger *slog.Logger
	server *echo.Echo
	db     *bun.DB
	worker *asynqwrapper.TaskClient
}

var (
	app     application
	appOnce sync.Once
)

func App() application {
	appOnce.Do(func() {
		configLogger()
		app.logger = initializers.InitLogger()
		app.db = initializers.InitDB()
		app.server = initializers.InitServer()
	})
	return app
}

func Script() application {
	appOnce.Do(func() {
		configLogger()
		app.logger = initializers.InitLogger()
		app.db = initializers.InitDB()
	})
	return app
}

func Test() application {
	appOnce.Do(func() {
		configLogger()
		app.logger = initializers.InitLogger()
	})
	return app
}

func AsynqWorker() application {
	appOnce.Do(func() {
		configLogger()
		app.logger = initializers.InitLogger()
		app.db = initializers.InitDB()
		worker, err := initializers.InitAsynq()
		if err != nil {
			panic(fmt.Sprintf("unable to start asynq worker %+v", err))
		}
		app.worker = worker
	})
	return app
}

func Server() *echo.Echo {
	return App().server
}

func DB() *bun.DB {
	return App().db
}

func L() *slog.Logger {
	return App().logger
}

func TaskClient() *asynqwrapper.TaskClient {
	return App().worker
}

func configLogger() {
	slog.Info("The app env is ", "env", config.Env())
	if config.IsDevelopment() {
		slog.Info("The app is running in development mode")
		slog.Info("The app", "settings", config.Settings())
		slog.Info("The app", "secrets", config.Secrets())
	}
}
