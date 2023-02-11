package main

import (
	"ledsuit/internal/config"
	"ledsuit/internal/input"
	"ledsuit/internal/led"
	"ledsuit/internal/pattern"
	"time"
)

func main() {

	suit := led.NewLedSuit(config.Selection)

	var patterns = [8]pattern.Pattern{
		pattern.NewDemoPattern(suit),
		pattern.NewLegsOnlyPattern(suit),
		pattern.NewBodyOnlyPattern(suit),
		pattern.NewTrafficPattern(suit),
		pattern.NewHueShiftPattern(suit),
		pattern.NewHAlternatePattern(suit),
		pattern.NewHeartPattern(suit),
		pattern.NewBlinkyPattern(suit),
	}

	current_pattern := 0
	controller := input.NewController()

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