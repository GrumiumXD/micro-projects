package pattern

import "ledsuit/internal/led"

type ArmsOnly struct {
	suit *led.LedSuit
	init bool
}

func NewArmsOnlyPattern(suit *led.LedSuit) Pattern {
	return &ArmsOnly{
		suit: suit,
		init: false,
	}
}

func (a *ArmsOnly) SetLEDs(tick uint32) {

	if !a.init {
		led.FillBuffer(a.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.LeftBody)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.LeftArm)[:], led.WHITE)
		led.FillBuffer(a.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.RightBody)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.RightArm)[:], led.WHITE)
		led.FillBuffer(a.suit.GetBuffer(led.Heart)[:], led.BLACK)

		a.init = true
	}
}

func (a *ArmsOnly) Reset() {

	a.init = false
}
