package timer

import (
	"jnsltk/look_away/internal/notifications"
	"time"
)

type Timer struct {
	TimerDuration time.Duration
	BreakDuration time.Duration
	Notifier *notifications.Notifier
}

func NewTimer(duration time.Duration, breakDuration time.Duration, notifier *notifications.Notifier) *Timer {
	return &Timer{duration, breakDuration, notifier}
}

func (t *Timer) Start() {
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
			t.Notifier.Notify(notificationMessages[i])
		}
	}
}
