package main

import (
	"ledsuit/internal/config"
	"ledsuit/internal/input"
	"ledsuit/internal/pattern"
	"ledsuit/internal/physical"
	"time"
)

func main() {

	suit := physical.NewLedSuit(config.Selection)

	var patterns = [8]pattern.Pattern{
		pattern.NewHellsBellsPattern(suit),
		pattern.NewDarkPattern(suit),
		pattern.NewBodyOnlyPattern(suit),
		pattern.NewTrafficPattern(suit),
		pattern.NewHueShiftPattern(suit),
		pattern.NewHAlternatePattern(suit),
		pattern.NewHeartPattern(suit),
		pattern.NewBlinkyPattern(suit),
	}

	current_pattern := 1
	controller := input.NewController(current_pattern)

	var tick uint32 = 0

	for {

		new_pattern := controller.GetValue()
		if new_pattern != current_pattern {
			patterns[current_pattern].Reset()
			current_pattern = new_pattern
		}

		patterns[current_pattern].SetLEDs(tick)
		suit.Display()

		tick++
		time.Sleep(time.Millisecond * 50)
	}
}
