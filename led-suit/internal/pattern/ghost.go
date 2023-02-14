package pattern

import "ledsuit/internal/led"

type Ghost struct {
	suit led.LedSuit
	init bool
}

func NewGhostPattern(suit led.LedSuit) Pattern {
	return &Ghost{
		suit: suit,
		init: false,
	}
}

func (g *Ghost) SetLEDs(tick uint32) {

	count := g.suit.GetCount()

	if !g.init {
		// everything not part of the heart green
		led.FillBuffer(g.suit.GetBuffer(led.LeftLeg)[:], led.GREEN)
		led.FillBuffer(g.suit.GetBuffer(led.LeftArm)[:], led.GREEN)
		led.FillBuffer(g.suit.GetBuffer(led.RightLeg)[:], led.GREEN)
		led.FillBuffer(g.suit.GetBuffer(led.RightArm)[:], led.GREEN)

		// body part before the heart
		led.FillBuffer(g.suit.GetBuffer(led.LeftBody)[:count.HBStart], led.GREEN)
		led.FillBuffer(g.suit.GetBuffer(led.RightBody)[:count.HBStart], led.GREEN)

		// body part after the heart
		led.FillBuffer(g.suit.GetBuffer(led.LeftBody)[count.HBEnd+1:], led.GREEN)
		led.FillBuffer(g.suit.GetBuffer(led.RightBody)[count.HBEnd+1:], led.GREEN)

		// fill the heart itself CYAN
		led.FillBuffer(g.suit.GetBuffer(led.LeftBody)[count.HBStart:count.HBEnd+1], led.CYAN)
		led.FillBuffer(g.suit.GetBuffer(led.RightBody)[count.HBStart:count.HBEnd+1], led.CYAN)
		led.FillBuffer(g.suit.GetBuffer(led.Heart)[:], led.CYAN)

		g.init = true
	}
}

func (g *Ghost) Reset() {

	g.init = false
}
