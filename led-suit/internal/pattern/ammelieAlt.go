package pattern

import (
	"image/color"
	"ledsuit/internal/led"
)

type AmmelieAlt struct {
	suit      led.LedSuit
	init      bool
	col       color.RGBA
	fadeIn    bool
	shift     int
	brighness uint32
}

const aCycle = 34

func NewAmmelieAltPattern(suit led.LedSuit, col color.RGBA) Pattern {
	return &AmmelieAlt{
		suit:      suit,
		init:      false,
		col:       col,
		fadeIn:    true,
		shift:     0,
		brighness: 0,
	}
}

func (a *AmmelieAlt) SetLEDs(brighness uint32) {

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

	c := led.Dim(a.col, float64(a.brighness)/float64(aCycle))

	for i := 0; i < int(a.suit.GetCount().Arm); i++ {
		if (i+a.shift)%2 == 0 {
			a.suit.GetBuffer(led.LeftArm)[i] = c
			a.suit.GetBuffer(led.RightArm)[i] = c
		} else {
			a.suit.GetBuffer(led.LeftArm)[i] = led.BLACK
			a.suit.GetBuffer(led.RightArm)[i] = led.BLACK
		}

	}

	if a.fadeIn {
		a.brighness++
		if a.brighness == aCycle {
			a.fadeIn = false
		}
	} else {
		a.brighness--
		if a.brighness == 0 {
			a.fadeIn = true
			a.shift = (a.shift + 1) % 2
		}
	}
}

func (a *AmmelieAlt) Reset() {

	a.init = false
	a.shift = 0
	a.fadeIn = true
	a.brighness = 0
}
