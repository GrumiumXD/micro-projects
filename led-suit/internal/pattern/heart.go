package pattern

import (
	"ledsuit/internal/led"
)

type Heart struct {
	suit led.LedSuit
	init bool
	on   bool
	tick uint32
}

func NewHeartPattern(suit led.LedSuit) Pattern {
	return &Heart{
		suit: suit,
		init: false,
		on:   true,
		tick: 0,
	}
}

func (h *Heart) SetLEDs(tick uint32) {

	count := h.suit.GetCount()

	if !h.init {
		// turn of ereything not part of the heart
		led.FillBuffer(h.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
		led.FillBuffer(h.suit.GetBuffer(led.LeftArm)[:], led.BLACK)
		led.FillBuffer(h.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
		led.FillBuffer(h.suit.GetBuffer(led.RightArm)[:], led.BLACK)

		// body part before the heart
		led.FillBuffer(h.suit.GetBuffer(led.LeftBody)[:count.HBStart], led.BLACK)
		led.FillBuffer(h.suit.GetBuffer(led.RightBody)[:count.HBStart], led.BLACK)

		// body part after the heart
		led.FillBuffer(h.suit.GetBuffer(led.LeftBody)[count.HBEnd+1:], led.BLACK)
		led.FillBuffer(h.suit.GetBuffer(led.RightBody)[count.HBEnd+1:], led.BLACK)

		h.init = true
	}

	if h.tick%4 == 0 {
		if h.on {
			// fill the heart parts
			led.FillBuffer(h.suit.GetBuffer(led.LeftBody)[count.HBStart:count.HBEnd+1], led.MAGENTA)
			led.FillBuffer(h.suit.GetBuffer(led.RightBody)[count.HBStart:count.HBEnd+1], led.MAGENTA)
			led.FillBuffer(h.suit.GetBuffer(led.Heart)[:], led.MAGENTA)
		} else {
			// fill the heart parts
			led.FillBuffer(h.suit.GetBuffer(led.LeftBody)[count.HBStart:count.HBEnd+1], led.BLACK)
			led.FillBuffer(h.suit.GetBuffer(led.RightBody)[count.HBStart:count.HBEnd+1], led.BLACK)
			led.FillBuffer(h.suit.GetBuffer(led.Heart)[:], led.BLACK)
		}

		h.on = !h.on

	}

	h.tick++
}

func (h *Heart) Reset() {

	h.tick = 0
	h.on = true
	h.init = false
}
