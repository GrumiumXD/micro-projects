package pattern

import (
	"image/color"
	"ledsuit/internal/led"
)

type HellsBells struct {
	suit   *led.LedSuit
	init   bool
	time   uint32
	fadeIn bool
}

const FADE_IN_TICKS = 6
const FADE_OUT_TICKS = 60

func NewHellsBellsPattern(suit *led.LedSuit) Pattern {
	return &HellsBells{
		suit:   suit,
		init:   false,
		time:   0,
		fadeIn: true,
	}
}

func (h *HellsBells) SetLEDs(tick uint32) {

	if !h.init {
		led.FillBuffer(h.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
		led.FillBuffer(h.suit.GetBuffer(led.LeftBody)[:], led.BLACK)
		led.FillBuffer(h.suit.GetBuffer(led.LeftArm)[:], led.BLACK)
		led.FillBuffer(h.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
		led.FillBuffer(h.suit.GetBuffer(led.RightBody)[:], led.BLACK)
		led.FillBuffer(h.suit.GetBuffer(led.RightArm)[:], led.BLACK)
		led.FillBuffer(h.suit.GetBuffer(led.Heart)[:], led.BLACK)

		h.init = true
	}

	var c color.RGBA

	if h.fadeIn {
		percent := float64(h.time) * 1.0 / FADE_IN_TICKS
		c = color.RGBA{uint8(255 * led.BRIGHTNESS * percent), 0, 0, 255}
		if h.time >= FADE_IN_TICKS {
			h.fadeIn = false
		}
	} else {
		if h.time >= FADE_IN_TICKS+FADE_OUT_TICKS {
			c = led.BLACK
		} else {
			percent := 1 - float64(h.time-FADE_IN_TICKS)*1.0/FADE_OUT_TICKS
			c = color.RGBA{uint8(255 * led.BRIGHTNESS * percent), 0, 0, 255}
		}
	}

	led.FillBuffer(h.suit.GetBuffer(led.LeftLeg)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.LeftBody)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.LeftArm)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.RightLeg)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.RightBody)[:], c)
	led.FillBuffer(h.suit.GetBuffer(led.RightArm)[:], c)

	h.time += 1
}

func (h *HellsBells) Reset() {

	h.init = false
	h.time = 0
	h.fadeIn = true
}
