package input

import (
	"machine"

	"time"
)

const (
	pin_mod = machine.GP20
	pin_0   = machine.GP16
	pin_1   = machine.GP17
	pin_2   = machine.GP18
	pin_3   = machine.GP19
)

type Controller struct {
	value   int
	buttons [4]machine.Pin
	mod     machine.Pin
	last    time.Time
}

func NewController(startValue int) *Controller {
	buttons := [4]machine.Pin{pin_0, pin_1, pin_2, pin_3}

	mod := pin_mod

	ctrl := &Controller{
		buttons: buttons,
		mod:     mod,
		value:   startValue,
		last:    time.Now(),
	}

	mod.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for i := 0; i < 4; i++ {
		buttons[i].Configure(machine.PinConfig{Mode: machine.PinInputPullup})

		buttons[i].SetInterrupt(machine.PinFalling, func(p machine.Pin) {
			ctrl.interrupt()
		})
	}

	return ctrl
}

func (c *Controller) interrupt() {
	if time.Since(c.last) < (200 * time.Millisecond) {
		return
	}

	// time.Sleep(10 * time.Millisecond)

	new_val := -1

	for i := 0; i < 4; i++ {
		if !c.buttons[i].Get() {
			new_val = i
			break
		}
	}

	if new_val == -1 {
		return
	}

	if !c.mod.Get() {
		new_val += 4
	}

	if new_val != c.value {
		c.value = new_val
		c.last = time.Now()
	}
}

func (c *Controller) GetValue() int {
	return c.value
}
