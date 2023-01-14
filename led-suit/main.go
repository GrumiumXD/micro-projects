package main

import (
	"ledsuit/input"
	"ledsuit/led"
	"ledsuit/pattern"
	"time"
)

func main() {

	suit := led.NewLedSuit(led.GR)
	// suit := led.NewLedSuit(led.RS)
	// suit := led.NewLedSuit(led.BK)
	// suit := led.NewLedSuit(led.GB)

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
