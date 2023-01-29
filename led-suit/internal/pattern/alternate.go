package pattern

import (
	"image/color"
	"ledsuit/internal/led"
)

type Alternate struct {
	suit   *led.LedSuit
	pos    int8
	left   bool
	init   bool
	length int8
}

func NewHAlternatePattern(suit *led.LedSuit) Pattern {
	c := suit.GetCount()

	length := c.Leg + c.Body + c.Arm

	return &Alternate{
		suit:   suit,
		pos:    0,
		left:   true,
		init:   false,
		length: length,
	}
}

func (a *Alternate) SetLEDs(tick uint32) {

	if !a.init {
		// clear everything during initilization
		led.FillBuffer(a.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.LeftBody)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.LeftArm)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.RightBody)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.RightArm)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.Heart)[:], led.BLACK)

		a.init = true
	}

	// update erery 4 ticks
	if tick%4 == 0 {
		if a.pos == a.length {
			// switch side when reaching the end
			a.pos = 0

			// reset the used strips
			if a.left {

				led.FillBuffer(a.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
				led.FillBuffer(a.suit.GetBuffer(led.LeftBody)[:], led.BLACK)
				led.FillBuffer(a.suit.GetBuffer(led.RightArm)[:], led.BLACK)
			} else {
				led.FillBuffer(a.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
				led.FillBuffer(a.suit.GetBuffer(led.RightBody)[:], led.BLACK)
				led.FillBuffer(a.suit.GetBuffer(led.LeftArm)[:], led.BLACK)

			}

			a.left = !a.left
		}

		r, g, b := led.Hsv2Rgb(6.0*float64(a.pos), 1.0, 1.0)
		c := color.RGBA{
			r, g, b, 255,
		}

		count := a.suit.GetCount()

		// leg
		if a.pos < count.Leg {
			index := count.Leg - (1 + a.pos)
			if a.left {
				a.suit.GetBuffer(led.LeftLeg)[index] = c
			} else {
				a.suit.GetBuffer(led.RightLeg)[index] = c
			}
		}

		// body
		if a.pos >= count.Leg && a.pos < count.Leg+count.Body {
			index := a.pos - count.Leg
			if a.left {
				a.suit.GetBuffer(led.LeftBody)[index] = c
			} else {
				a.suit.GetBuffer(led.RightBody)[index] = c
			}
		}

		// arm
		if a.pos >= count.Leg+count.Body && a.pos < count.Leg+count.Body+count.Arm {
			index := a.pos - (count.Leg + count.Body)
			if a.left {
				a.suit.GetBuffer(led.RightArm)[index] = c
			} else {
				a.suit.GetBuffer(led.LeftArm)[index] = c
			}
		}

		// advance the position
		a.pos = a.pos + 1
	}
}

func (a *Alternate) Reset() {

	a.left = true
	a.pos = 0
	a.init = false
}
