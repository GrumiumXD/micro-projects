package main

import (
	"ledtest/input"
	"machine"
	"time"
)

// const LED_COUNT = 9

var tick uint32 = 0

// var CP = 0
// var last time.Time = time.Now()

var patterns = []PatternEx{
	&Different{},
	&Arms{},
	&Legs{},
	&Body{},
	&Hueshift{},
	&Switching{},
	&Heart{},
	&Blinky{},
}

// func switchPattern(index int) {
// 	if time.Since(last) < (200 * time.Millisecond) {
// 		return
// 	}
// 	last = time.Now()

// 	newIndex := index
// 	if !PIN_MOD.Get() {
// 		// mod key pressed
// 		newIndex += 4
// 	}

// 	if newIndex != CP {
// 		patterns[CP].Reset()
// 		CP = newIndex
// 	}
// }

func main() {
	// patterns := []Pattern{
	// 	&SolidRed{},
	// 	&Selective{},
	// 	&Blinky{},
	// 	&HueShift{},
	// 	&Snake{},
	// }

	// switch pattern button setup
	// switchPin0 := machine.GP16
	// switchPin0.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	// switchPin0.SetInterrupt(machine.PinFalling, func(p machine.Pin) {
	// 	switchPattern(0)
	// })
	// switchPin1 := machine.GP17
	// switchPin1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	// switchPin1.SetInterrupt(machine.PinFalling, func(p machine.Pin) {
	// 	switchPattern(1)
	// })
	// switchPin2 := machine.GP18
	// switchPin2.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	// switchPin2.SetInterrupt(machine.PinFalling, func(p machine.Pin) {
	// 	switchPattern(2)
	// })
	// switchPin3 := machine.GP19
	// switchPin3.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	// switchPin3.SetInterrupt(machine.PinFalling, func(p machine.Pin) {
	// 	switchPattern(3)
	// })
	// switchPin4 := machine.GP20
	// switchPin4.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	// switchPin4.SetInterrupt(machine.PinFalling, func(p machine.Pin) {
	// 	switchPattern(4)
	// })

	// ws2812 setup
	// ledPin := machine.GP15
	// ledPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// strip := ws2812.New(ledPin)

	// buf := make([]color.RGBA, LED_COUNT)

	strips, buffers := stripSetup()
	// inputSetup(switchPattern)

	current_pattern := 0
	ctrl := input.NewController([4]machine.Pin{machine.GP16, machine.GP17, machine.GP18, machine.GP19}, machine.GP20)

	for {

		// patterns[CP].SetLEDs(buf, tick)
		// strip.WriteColors(buf)

		new_pattern := ctrl.GetValue()
		if new_pattern != current_pattern {
			patterns[current_pattern].Reset()
			current_pattern = new_pattern
		}

		patterns[current_pattern].SetLEDs(buffers, tick)
		for i := 0; i < STRIP_COUNT; i++ {
			strips[i].WriteColors(buffers[i])
		}

		tick++
		time.Sleep(time.Millisecond * 50)
	}
}
