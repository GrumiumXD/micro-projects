package pattern

import (
	"image/color"
	"ledsuit/internal/led"
)

type Cannon struct {
	suit   led.LedSuit
	init   bool
	pos    int
	colPos int
	left   bool
}

var beamColors = [2]color.RGBA{
	led.MAGENTA,
	led.YELLOW,
}

const beamLength = 3

func NewCannonPattern(suit led.LedSuit, left bool) Pattern {
	return &Cannon{
		suit:   suit,
		init:   false,
		colPos: 0,
		pos:    -beamLength,
		left:   left,
	}
}

func (c *Cannon) SetLEDs(tick uint32) {

	if !c.init {
		led.FillBuffer(c.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
		led.FillBuffer(c.suit.GetBuffer(led.LeftBody)[:], led.BLACK)
		led.FillBuffer(c.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
		led.FillBuffer(c.suit.GetBuffer(led.RightBody)[:], led.BLACK)
		led.FillBuffer(c.suit.GetBuffer(led.RightArm)[:], led.BLACK)

		led.FillBuffer(c.suit.GetBuffer(led.Heart)[:], led.CYAN)
		if c.left {
			led.FillBuffer(c.suit.GetBuffer(led.LeftArm)[:], led.CYAN)
			led.FillBuffer(c.suit.GetBuffer(led.RightArm)[:], led.BLACK)
		} else {
			led.FillBuffer(c.suit.GetBuffer(led.LeftArm)[:], led.BLACK)
			led.FillBuffer(c.suit.GetBuffer(led.RightArm)[:], led.CYAN)
		}

		c.init = true
	}

	// arms
	for i := 0; i < int(c.suit.GetCount().Arm); i++ {
		if i >= c.pos && i < c.pos+beamLength {
			if c.left {
				c.suit.GetBuffer(led.RightArm)[i] = beamColors[c.colPos]
			} else {
				c.suit.GetBuffer(led.LeftArm)[i] = beamColors[c.colPos]
			}

		} else {
			if c.left {
				c.suit.GetBuffer(led.RightArm)[i] = led.BLACK
			} else {
				c.suit.GetBuffer(led.LeftArm)[i] = led.BLACK
			}
		}
	}

	c.pos += 2
	if c.pos >= int(c.suit.GetCount().Arm) {
		c.pos = -beamLength
		c.colPos = (c.colPos + 1) % len(beamColors)
	}

}

func (c *Cannon) Reset() {

	c.init = false
	c.colPos = 0
	c.pos = -beamLength
}
