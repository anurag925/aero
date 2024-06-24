package asynqworkers

import (
	"context"

	"github.com/anurag925/aero/core/logger"
	"github.com/hibiken/asynq"
)

type asyncProcessWorker struct {
}

func NewAsyncProcessWorker() *asyncProcessWorker {
	return &asyncProcessWorker{}
}

func (o *asyncProcessWorker) ProcessTask(ctx context.Context, t *asynq.Task) error {
	logger.Info(ctx, "running the async worker")
	return nil
}
