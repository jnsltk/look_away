package main

import (
	"jnsltk/look_away/internal/timer"
	"time"
)

func main() {
	t := timer.NewTimer(2 * time.Second, time.Second)
	// t := timer.NewTimer(20 * time.Minue, 20 * time.Second)
	t.Start()
}