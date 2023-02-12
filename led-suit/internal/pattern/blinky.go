package pattern

import (
	"image/color"
	"ledsuit/internal/led"
)

type Blinky struct {
	suit led.LedSuit
	init bool
}

func NewBlinkyPattern(suit led.LedSuit) Pattern {
	return &Blinky{
		suit: suit,
		init: false,
	}
}

func (b *Blinky) SetLEDs(tick uint32) {

	if !b.init {
		led.FillBuffer(b.suit.GetBuffer(led.Heart)[:], led.BLACK)

		b.init = true
	}

	var c color.RGBA
	if tick%2 == 0 {
		c = led.WHITE
	} else {
		c = led.BLACK
	}

	// set strips
	led.FillBuffer(b.suit.GetBuffer(led.LeftLeg)[:], c)
	led.FillBuffer(b.suit.GetBuffer(led.LeftBody)[:], c)
	led.FillBuffer(b.suit.GetBuffer(led.LeftArm)[:], c)
	led.FillBuffer(b.suit.GetBuffer(led.RightLeg)[:], c)
	led.FillBuffer(b.suit.GetBuffer(led.RightBody)[:], c)
	led.FillBuffer(b.suit.GetBuffer(led.RightArm)[:], c)
}

func (b *Blinky) Reset() {

	b.init = false
}
