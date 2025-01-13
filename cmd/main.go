package main

import (
	"fmt"
	"jnsltk/look_away/internal/config"
	"jnsltk/look_away/internal/notifications"
	"jnsltk/look_away/internal/timer"
)

func main() {
	cfg, err := config.LoadConfig("internal/config/config.yaml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	notifier := notifications.NewNotifier(cfg.Notifications)
	t := timer.NewTimer(cfg.GetTimerDuration(), cfg.GetBreakSeconds(), notifier)
	t.Start()
}