package asynqworkers

import (
	"context"

	"github.com/anurag925/aero/core/logger"
	"github.com/hibiken/asynq"
)

type asyncCronWorker struct {
}

func NewAsyncCronWorker() *asyncCronWorker {
	return &asyncCronWorker{}
}

func (o *asyncCronWorker) ProcessTask(ctx context.Context, t *asynq.Task) error {
	logger.Info(ctx, "running the async cron worker")
	return nil
}
