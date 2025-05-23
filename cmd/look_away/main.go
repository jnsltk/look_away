package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"jnsltk/look_away/internal/config"
	"jnsltk/look_away/internal/notifications"
	"jnsltk/look_away/internal/timer"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"gopkg.in/yaml.v3"
)

const APP_NAME string = "look_away"
const CONFIG_FILE_NAME string = "config.yaml"

func main() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("User config directory not found", err)
	}

	configPath := filepath.Join(userConfigDir, APP_NAME, CONFIG_FILE_NAME)

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Config file not found. Creating default config...")
		err := createDefaultConfig(configPath)
		if err != nil {
			log.Fatalf("Error creating default config", err)
		}
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	var (
		customConfigPath    string
		customTimerDuration int
		customBreakDuration int
		showConfigLocation  bool
		printConfig         bool
		help                bool
	)

	flag.StringVar(&customConfigPath, "config", "$HOME/.config/look_away/config.yml", "Path to the yaml config file")
	flag.IntVar(&customTimerDuration, "duration", 0, "Timer duration in minutes (overrides config)")
	flag.IntVar(&customBreakDuration, "break-duration", 0, "Break duration in seconds (overrides config)")
	flag.BoolVar(&showConfigLocation, "config-path", false, "Print default config location")
	flag.BoolVar(&printConfig, "print-config", false, "Print the current configuration")
	flag.BoolVar(&help, "help", false, "Show help message")
	flag.BoolVar(&help, "h", false, "Show help message")

	flag.Parse()

	if help {
		printHelp()
		return
	}

	if showConfigLocation {
		fmt.Println(configPath)
		return
	}

	if printConfig {
		printCurrentConfig(cfg)
		return
	}

	if customTimerDuration > 0 {
		cfg.Timer.DurationMinutes = customTimerDuration
	}
	if customBreakDuration > 0 {
		cfg.Timer.BreakSeconds = customBreakDuration
	}

	notifier := notifications.NewNotifier(cfg.Notifications)
	t := timer.NewTimer(cfg.GetTimerDuration(), cfg.GetBreakSeconds(), notifier)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("20-20-20 timer started!")
	fmt.Println("Press Ctrl+C to quit.")

	go t.Start(ctx)

	<-quitChan

	fmt.Println("\nQuitting application...")
}

func createDefaultConfig(configPath string) error {
	defaultConfig := config.AppConfig{
		Timer: config.TimerConfig{
			DurationMinutes: 20,
			BreakSeconds:    20,
		},
		Notifications: config.NotificationConfig{
			UseAlert: true,
		},
	}

	data, err := yaml.Marshal(defaultConfig)
	if err != nil {
		return err
	}

	configDir := filepath.Dir(configPath)
	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func printHelp() {
	fmt.Println("look_away -- 20-20-20 timer app")
	fmt.Println("The 20-20-20 method is a simple way to minimise eyestrain. Every 20 minutes, " +
		"look up from your screen and focus on an item approximately 20 feet (~6m) away for at least 20 seconds.")
	fmt.Println("Usage:")
	fmt.Println("  --config          Path to the YAML config file (Optional, default: internal/config/config.yaml)")
	fmt.Println("  --duration        Timer duration in minutes (Optional, overrides config)")
	fmt.Println("  --break-duration  Break duration in seconds (Optional, overrides config)")
	fmt.Println("  --config-path     Print default config location")
	fmt.Println("  --print-config    Print the current configuration")
	fmt.Println("  --help, -h        Show this help message")
	fmt.Println("\nExample: look_away --duration=25 --break_duration=30")
}

func printCurrentConfig(cfg *config.AppConfig) {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatalf("Error printing config", err)
	}

	fmt.Println(string(data))
}
