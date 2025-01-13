package timer

import (
	"context"
	"fmt"
	"jnsltk/look_away/internal/notifications"
	"time"
)

type Timer struct {
	TimerDuration time.Duration
	BreakDuration time.Duration
	Notifier      *notifications.Notifier
}

func NewTimer(duration time.Duration, breakDuration time.Duration, notifier *notifications.Notifier) *Timer {
	return &Timer{
		duration,
		breakDuration,
		notifier,
	}
}

func (t *Timer) Start(ctx context.Context) {
	durations := []time.Duration{t.TimerDuration, t.BreakDuration}
	notificationMessages := []string{
		"Time to rest your eyes! Look at least 20 ft (~6m) away for at least 20 seconds!",
		"That's enough, go back to work!",
	}

	for {
		for i, duration := range durations {
			timer := time.NewTimer(duration)
			ticker := time.NewTicker(1 * time.Second)

			innerloop:
			for remaining := duration; remaining > 0; remaining -= 1 * time.Second {
				select {
				case <-ctx.Done():
					timer.Stop()
					ticker.Stop()
					fmt.Println("Timer stopped")
				case <-timer.C:
					t.Notifier.Notify(notificationMessages[i])
					break innerloop
				case <-ticker.C:
					minutes := int(remaining.Minutes())
					seconds := int(remaining.Seconds()) % 60
					fmt.Printf("\r%02d:%02d remaining", minutes, seconds)
				}
			}
			timer.Stop()
			ticker.Stop()
		}
	}
}
