package pattern

import (
	"image/color"
	"ledsuit/internal/led"
)

type HueShift struct {
	suit led.LedSuit
	hue  uint16
	init bool
}

func NewHueShiftPattern(suit led.LedSuit) Pattern {
	return &HueShift{
		suit: suit,
		hue:  0,
		init: false,
	}
}

func (h *HueShift) SetLEDs(tick uint32) {

	if !h.init {
		// shut of heart during initilization
		led.FillBuffer(h.suit.GetBuffer(led.Heart)[:], led.BLACK)

		h.init = true
	}

	// calculate color from hue
	r, g, b := led.Hsv2Rgb(float64(h.hue), 1.0, 1.0)
	c := color.RGBA{
		r, g, b, 255,
	}

	// set strips
	led.FillBuffer(h.suit.GetBuffer(led.LeftLeg)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.LeftBody)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.LeftArm)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.RightLeg)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.RightBody)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.RightArm)[:], c)

	// advance hue
	h.hue = (h.hue + 1) % 360
}

func (h *HueShift) Reset() {

	h.hue = 0
	h.init = false
}
