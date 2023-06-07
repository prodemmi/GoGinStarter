package scheduler

import (
	"github.com/go-co-op/gocron"
	"time"
)

type Schedule interface {
	Run() (*gocron.Scheduler, error)
}

type schedule struct{}

func (s *schedule) Run() (*gocron.Scheduler, error) {
	cron := gocron.NewScheduler(time.UTC)
	for _, task := range Tasks {
		_, err := task(cron)
		if err != nil {
			return nil, err
		}
	}
	cron.StartAsync()

	return cron, nil
}

func ProvideSchedule() Schedule {
	return &schedule{}
}
