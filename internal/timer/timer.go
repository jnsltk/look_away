package timer

import (
	"time"
	"github.com/gen2brain/beeep"
)

type Timer struct {
	TimerDuration      time.Duration
	BreakDuration time.Duration
}

func NewTimer(duration time.Duration, breakDuration time.Duration) *Timer {
	return &Timer{duration, breakDuration}
}

func (t Timer) Start() {
	durations := []time.Duration{t.TimerDuration, t.BreakDuration}
	notificationMessages := []string{
		"Time to rest your eyes! Look at least 20 ft (~6m) away for at least 20 seconds!", 
		"That's enough, go back to work!",
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
			beeep.Alert("look_away", notificationMessages[i], "")
		}
	}
}
