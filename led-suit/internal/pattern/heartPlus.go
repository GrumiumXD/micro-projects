package pattern

import (
	"image/color"
	"ledsuit/internal/led"
)

type HeartPlus struct {
	suit led.LedSuit
	init bool
	col  color.RGBA
}

func NewHeartPlusPattern(suit led.LedSuit, col color.RGBA) Pattern {
	return &HeartPlus{
		suit: suit,
		init: false,
		col:  col,
	}
}

func (h *HeartPlus) SetLEDs(tick uint32) {

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

		// fill the heart parts
		led.FillBuffer(h.suit.GetBuffer(led.LeftBody)[count.HBStart:count.HBEnd+1], led.MAGENTA)
		led.FillBuffer(h.suit.GetBuffer(led.RightBody)[count.HBStart:count.HBEnd+1], led.MAGENTA)
		led.FillBuffer(h.suit.GetBuffer(led.Heart)[:], led.MAGENTA)

		h.init = true
	}

	var c color.RGBA
	if tick%2 == 0 {
		c = h.col
	} else {
		c = led.BLACK
	}

	// change everything not part of the heart
	led.FillBuffer(h.suit.GetBuffer(led.LeftLeg)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.LeftArm)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.RightLeg)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.RightArm)[:], c)

	// body part before the heart
	led.FillBuffer(h.suit.GetBuffer(led.LeftBody)[:count.HBStart], c)
	led.FillBuffer(h.suit.GetBuffer(led.RightBody)[:count.HBStart], c)

	// body part after the heart
	led.FillBuffer(h.suit.GetBuffer(led.LeftBody)[count.HBEnd+1:], c)
	led.FillBuffer(h.suit.GetBuffer(led.RightBody)[count.HBEnd+1:], c)

}

func (h *HeartPlus) Reset() {
	h.init = false
}
