package main

import (
	"flag"
	"fmt"
	"jnsltk/look_away/internal/config"
	"jnsltk/look_away/internal/notifications"
	"jnsltk/look_away/internal/timer"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var configPath string
	var timerDuration int
	var breakDuration int
	var useAlert *bool
	var help bool

	flag.StringVar(&configPath, "config", "internal/config/config.yaml", "Path to the yaml config file")
	flag.IntVar(&timerDuration, "duration", 0, "Timer duration in minutes (overrides config)")
	flag.IntVar(&breakDuration, "break_duration", 0, "Break duration in seconds (overrides config)")
	useAlert = flag.Bool("alert", false, "Use alert instead of notification (overrides config)")
	flag.BoolVar(&help, "help", false, "Show help message")
	flag.BoolVar(&help, "h", false, "Show help message")

	flag.Parse()

	if help {
		fmt.Println("look_away -- 20-20-20 timer app")
		fmt.Println("The 20-20-20 method is a simple way to minimise eyestrain. Every 20 minutes, " + 
		"look up from your screen and focus on an item approximately 20 feet (~6m) away for at least 20 seconds.")
		fmt.Println("Usage:")
		fmt.Println("  -config          Path to the YAML config file (Optional, default: internal/config/config.yaml)")
        fmt.Println("  -duration        Timer duration in minutes (Optional, overrides config)")
        fmt.Println("  -break_duration  Break duration in seconds (Optional, overrides config)")
        fmt.Println("  -alert           Use alert instead of notification (Optional, overrides config)")
        fmt.Println("  -help            Show this help message")
		fmt.Println("\nExample: look_away -duration=25 -break_duration=30 -alert=false")
        return
	}

	cfg, err := config.LoadConfig("internal/config/config.yaml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	if timerDuration > 0 {
		cfg.Timer.DurationMinutes = timerDuration
	}
	if breakDuration > 0 {
		cfg.Timer.BreakSeconds = breakDuration
	}
	if useAlert != nil {
		cfg.Notifications.UseAlert = *useAlert
	}

	notifier := notifications.NewNotifier(cfg.Notifications)
	t := timer.NewTimer(cfg.GetTimerDuration(), cfg.GetBreakSeconds(), notifier)
	
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("20-20-20 timer started!")
	fmt.Println("Press Ctrl+C to quit.")

	go t.Start()

	<-quitChan
	fmt.Println("\nApp quit gracefully.")
}
