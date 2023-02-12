package pattern

import "ledsuit/internal/led"

type Demo struct {
	suit led.LedSuit
	init bool
}

func NewDemoPattern(suit led.LedSuit) Pattern {
	return &Demo{
		suit: suit,
		init: false,
	}
}

func (d *Demo) SetLEDs(tick uint32) {

	if !d.init {
		// initialize every strip with a different color
		led.FillBuffer(d.suit.GetBuffer(led.LeftLeg)[:], led.RED)
		led.FillBuffer(d.suit.GetBuffer(led.LeftBody)[:], led.BLUE)
		led.FillBuffer(d.suit.GetBuffer(led.LeftArm)[:], led.CYAN)
		led.FillBuffer(d.suit.GetBuffer(led.RightLeg)[:], led.GREEN)
		led.FillBuffer(d.suit.GetBuffer(led.RightBody)[:], led.YELLOW)
		led.FillBuffer(d.suit.GetBuffer(led.RightArm)[:], led.MAGENTA)
		led.FillBuffer(d.suit.GetBuffer(led.Heart)[:], led.WHITE)

		d.init = true
	}
}

func (d *Demo) Reset() {

	d.init = false
}
