package asynqworkers

import (
	"github.com/anurag925/aero/lib/utils/asynqwrapper"
)

type asynqWorker struct {
	w asynqwrapper.AsyncTask
}

func New(w asynqwrapper.AsyncTask) *asynqWorker {
	return &asynqWorker{w}
}

func (s *asynqWorker) ScheduleTask() error {
	if err := s.w.New("async_cron_worker").MaxRetries(5).Periodic("00 12 * * *").Save(); err != nil {
		return err
	}
	return nil
}

func (s *asynqWorker) RegisterTask() {
	s.w.AddTask("async_process_worker", NewAsyncProcessWorker())
	s.w.AddTask("async_cron_worker", NewAsyncCronWorker())
}

func (s *asynqWorker) StartServer() error {
	return s.w.StartServer()
}

func (s *asynqWorker) StartScheduler() error {
	return s.w.StartScheduler()
}
