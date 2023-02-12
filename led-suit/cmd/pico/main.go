package main

import (
	"ledsuit/internal/config"
	"ledsuit/internal/input"
	"ledsuit/internal/led"
	"ledsuit/internal/pattern"
	"ledsuit/internal/physical"
	"time"
)

func main() {

	suit := physical.NewLedSuit(config.Selection)

	var patterns [8]pattern.Pattern

	if config.Selection == "GR" {
		patterns = [8]pattern.Pattern{
			pattern.NewHellsBellsPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewAmmeliePattern(suit, led.CYAN),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
		}
	} else if config.Selection == "RS" {
		patterns = [8]pattern.Pattern{
			pattern.NewHellsBellsPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewAmmeliePattern(suit, led.ORANGE),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
		}
	} else if config.Selection == "BK" {
		patterns = [8]pattern.Pattern{
			pattern.NewHellsBellsPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewAmmeliePattern(suit, led.BRIGHT_GREEN),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
		}
	} else if config.Selection == "GB" {
		patterns = [8]pattern.Pattern{
			pattern.NewHellsBellsPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewAmmeliePattern(suit, led.YELLOW),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
			pattern.NewDarkPattern(suit),
		}
	}

	// start pattern
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
