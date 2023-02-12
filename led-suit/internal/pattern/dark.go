package pattern

import "ledsuit/internal/led"

type Dark struct {
	suit *led.LedSuit
	init bool
}

func NewDarkPattern(suit *led.LedSuit) Pattern {
	return &Dark{
		suit: suit,
		init: false,
	}
}

func (a *Dark) SetLEDs(tick uint32) {

	if !a.init {
		led.FillBuffer(a.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.LeftBody)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.LeftArm)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.RightBody)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.RightArm)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.Heart)[:], led.BLACK)

		a.init = true
	}
}

func (a *Dark) Reset() {

	a.init = false
}
