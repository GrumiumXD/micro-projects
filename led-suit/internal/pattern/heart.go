package pattern

import (
	"image/color"
	"ledsuit/internal/led"
)

type Heart struct {
	suit    *led.LedSuit
	timer   uint8
	forward bool
	init    bool
}

func NewHeartPattern(suit *led.LedSuit) Pattern {
	return &Heart{
		suit:    suit,
		timer:   0,
		forward: false,
		init:    false,
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

	var hue float64
	if h.forward {
		hue = 300.0 + float64(h.timer)
	} else {
		hue = 360.0 - float64(h.timer)
	}

	r, g, b := led.Hsv2Rgb(hue, 1.0, 1.0)
	c := color.RGBA{
		r, g, b, 255,
	}

	// fill the heart parts
	led.FillBuffer(h.suit.GetBuffer(led.LeftBody)[count.HBStart:count.HBEnd+1], c)
	led.FillBuffer(h.suit.GetBuffer(led.RightBody)[count.HBStart:count.HBEnd+1], c)
	led.FillBuffer(h.suit.GetBuffer(led.Heart)[:], c)

	h.timer += 1

	if h.timer == 60 {
		h.forward = !h.forward
		h.timer = 0
	}
}

func (h *Heart) Reset() {

	h.timer = 0
	h.forward = false
	h.init = false
}
