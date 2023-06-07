package schedules

import (
	"fmt"
	"github.com/go-co-op/gocron"
)

func ExampleTask(scheduler *gocron.Scheduler) (*gocron.Job, error) {
	return scheduler.Every("5s").Do(func() {
		fmt.Println("RUNNNNNN")
	})
}
