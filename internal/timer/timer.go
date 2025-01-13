package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	Duration      time.Duration
	BreakDuration time.Duration
}

func NewTimer(duration time.Duration, breakDuration time.Duration) *Timer {
	return &Timer{duration, breakDuration}
}

func (t Timer) Start() {
	durations := []time.Duration{t.Duration, t.BreakDuration}
	functions := []func(){
		func() { fmt.Println("done") }, 
		func() { fmt.Println("break done") },
	}

	var timer *time.Timer
	defer func() {
		if timer != nil {
			timer.Stop()
		}
	}()

	for {
		for i, duration := range durations {
			timer = time.NewTimer(duration)
			<-timer.C
			functions[i]()
		}
	}
}
