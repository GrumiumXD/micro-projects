package pattern

import (
	"image/color"
	"ledsuit/internal/led"
)

var tcolors = []color.RGBA{
	led.GREEN,
	led.RED,
	led.CYAN,
	led.YELLOW,
	led.BLUE,
	led.BRIGHT_GREEN,
	led.MAGENTA,
	led.ORANGE,
	led.WHITE,
}

type Tetris struct {
	suit  led.LedSuit
	init  bool
	tick  uint32
	start int
	mode  int
}

func NewTetrisPattern(suit led.LedSuit, start int) Pattern {
	return &Tetris{
		suit:  suit,
		init:  false,
		tick:  0,
		start: start,
		mode:  start,
	}
}

func (t *Tetris) SetLEDs(tick uint32) {
	m := len(tcolors)

	if !t.init {
		// initialize every strip with a different color
		led.FillBuffer(t.suit.GetBuffer(led.Heart)[:], led.BLACK)

		t.init = true
	}

	if t.tick%8 == 0 {
		led.FillBuffer(t.suit.GetBuffer(led.LeftLeg)[:], tcolors[t.mode])
		led.FillBuffer(t.suit.GetBuffer(led.LeftBody)[:], tcolors[(t.mode+1)%m])
		led.FillBuffer(t.suit.GetBuffer(led.LeftArm)[:], tcolors[(t.mode+2)%m])
		led.FillBuffer(t.suit.GetBuffer(led.RightLeg)[:], tcolors[(t.mode+3)%m])
		led.FillBuffer(t.suit.GetBuffer(led.RightBody)[:], tcolors[(t.mode+4)%m])
		led.FillBuffer(t.suit.GetBuffer(led.RightArm)[:], tcolors[(t.mode+5)%m])

		t.mode = (t.mode + 1) % m
	}

	t.tick++
}

func (t *Tetris) Reset() {

	t.init = false
	t.mode = t.start
	t.tick = 0
}
