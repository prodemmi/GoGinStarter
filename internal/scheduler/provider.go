package scheduler

import (
	"github.com/go-co-op/gocron"
)

var Tasks = []func(scheduler *gocron.Scheduler) (*gocron.Job, error){
	//schedules.ExampleTask,
}
