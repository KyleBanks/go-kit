// Package job provides the ability to execute timed jobs in their own goroutine.
package job

import "time"

type Job interface {
	// Run is called when the job is triggered.
	Run()

	// SleepTime returns the amount of time to sleep before running
	// the job again.
	SleepTime() time.Duration
}

// RegisterJob schedules a job for execution
func RegisterJob(j Job) {
	go func(j Job) {
		for {
			j.Run()
			time.Sleep(j.SleepTime())
		}
	}(j)
}
