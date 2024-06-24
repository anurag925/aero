package main

import (
	"log/slog"

	"github.com/anurag925/aero/app/workers/asynqworkers"
	"github.com/anurag925/aero/core"
)

func main() {
	core.AsynqWorker()
	workers := asynqworkers.New(core.TaskClient())
	workers.RegisterTask()
	if err := workers.ScheduleTask(); err != nil {
		slog.Error("unable to schedule application task", err)
		panic("unable to schedule application task")
	}

	go workers.StartScheduler()
	if err := workers.StartServer(); err != nil {
		slog.Error("unable to start the async task with error %+v", err)
		panic("unable to start the async server")
	}
}
