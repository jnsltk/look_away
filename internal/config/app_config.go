package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type TimerConfig struct {
	DurationMinutes int `yaml:"duration_minutes"`
	BreakSeconds int `yaml:"break_seconds"`
}

type NotificationConfig struct {
	UseAlert bool `yaml:"use_alert"`
}

type AppConfig struct {
	Timer TimerConfig `yaml:"timer"`
	Notifications NotificationConfig `yaml:"notifications"`
}

func LoadConfig(path string) (*AppConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read config file %v", err)
	}

	var config AppConfig
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("failed to parse yaml %v", err)
	}

	return &config, nil
}

func (c *AppConfig) GetTimerDuration() time.Duration {
	return time.Duration(c.Timer.DurationMinutes) * time.Minute
}

func (c *AppConfig) GetBreakSeconds() time.Duration {
	return time.Duration(c.Timer.BreakSeconds) * time.Second
}