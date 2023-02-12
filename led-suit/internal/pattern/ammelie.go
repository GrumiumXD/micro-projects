package pattern

import (
	"image/color"
	"ledsuit/internal/led"
)

type Ammelie struct {
	suit led.LedSuit
	init bool
	col  color.RGBA
	mode int
	tick uint32
}

var numbers = []int{1, 2, 3, 5, 7, 5, 3, 2}

func NewAmmeliePattern(suit led.LedSuit, col color.RGBA) Pattern {
	return &Ammelie{
		suit: suit,
		init: false,
		col:  col,
		tick: 0,
		mode: 0,
	}
}

func (a *Ammelie) SetLEDs(tick uint32) {

	if !a.init {
		led.FillBuffer(a.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.LeftBody)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.LeftArm)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.RightBody)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.RightArm)[:], led.BLACK)
		led.FillBuffer(a.suit.GetBuffer(led.Heart)[:], led.BLACK)

		a.init = true
	}

	if a.tick%17 == 0 {
		a.mode = (a.mode + 1) % 8

		for i := 0; i < int(a.suit.GetCount().Arm); i++ {
			if (i+a.mode)%numbers[a.mode] == 0 {
				a.suit.GetBuffer(led.LeftArm)[i] = a.col
				a.suit.GetBuffer(led.RightArm)[i] = a.col
			} else {
				a.suit.GetBuffer(led.LeftArm)[i] = led.BLACK
				a.suit.GetBuffer(led.RightArm)[i] = led.BLACK
			}
		}
	}

	a.tick++
}

func (a *Ammelie) Reset() {

	a.init = false
	a.mode = 0
	a.tick = 0
}
