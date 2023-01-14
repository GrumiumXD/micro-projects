package pattern

import "ledsuit/led"

type BodyOnly struct {
	suit *led.LedSuit
	init bool
}

func NewBodyOnlyPattern(suit *led.LedSuit) Pattern {
	return &BodyOnly{
		suit: suit,
		init: false,
	}
}

func (b *BodyOnly) SetLEDs(tick uint32) {

	if !b.init {
		led.FillBuffer(b.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
		led.FillBuffer(b.suit.GetBuffer(led.LeftBody)[:], led.GREEN)
		led.FillBuffer(b.suit.GetBuffer(led.LeftArm)[:], led.BLACK)
		led.FillBuffer(b.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
		led.FillBuffer(b.suit.GetBuffer(led.RightBody)[:], led.GREEN)
		led.FillBuffer(b.suit.GetBuffer(led.RightArm)[:], led.BLACK)
		led.FillBuffer(b.suit.GetBuffer(led.Heart)[:], led.BLACK)

		b.init = true
	}
}

func (b *BodyOnly) Reset() {

	b.init = false
}
