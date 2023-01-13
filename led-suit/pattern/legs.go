package pattern

import "ledsuit/led"

type LegsOnly struct {
	suit *led.LedSuit
	init bool
}

func NewLegsOnlyPattern(suit *led.LedSuit) Pattern {
	return &LegsOnly{
		suit: suit,
		init: false,
	}
}

func (l *LegsOnly) SetLEDs(tick uint32) {

	if !l.init {
		led.FillBuffer(l.suit.GetBuffer(led.LeftLeg)[:], led.WHITE)
		led.FillBuffer(l.suit.GetBuffer(led.LeftBody)[:], led.BLACK)
		led.FillBuffer(l.suit.GetBuffer(led.LeftArm)[:], led.BLACK)
		led.FillBuffer(l.suit.GetBuffer(led.RightLeg)[:], led.WHITE)
		led.FillBuffer(l.suit.GetBuffer(led.RightBody)[:], led.BLACK)
		led.FillBuffer(l.suit.GetBuffer(led.RightArm)[:], led.BLACK)
		led.FillBuffer(l.suit.GetBuffer(led.Heart)[:], led.BLACK)

		l.init = true
	}
}

func (l *LegsOnly) Reset() {

	l.init = false
}
