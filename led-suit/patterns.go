package main

import (
	"image/color"
)

type Pattern interface {
	SetLEDs(buf []color.RGBA, tick uint32)
}

type PatternEx interface {
	SetLEDs(buf [][]color.RGBA, tick uint32)
	Reset()
}

// 0
type Different struct {
	init bool
}

func (p *Different) SetLEDs(buf [][]color.RGBA, tick uint32) {

	if !p.init {
		fill(buf[LB][:], RED)
		fill(buf[RB][:], GREEN)
		fill(buf[LK][:], BLUE)
		fill(buf[RK][:], YELLOW)
		fill(buf[LA][:], CYAN)
		fill(buf[RA][:], MAGENTA)
		fill(buf[H][:], WHITE)

		p.init = true
	}
}

func (p *Different) Reset() {

	p.init = false
}

// 1
type Arms struct {
	init bool
}

func (p *Arms) SetLEDs(buf [][]color.RGBA, tick uint32) {

	if !p.init {
		fill(buf[LB][:], BLACK)
		fill(buf[RB][:], BLACK)
		fill(buf[LK][:], BLACK)
		fill(buf[RK][:], BLACK)
		fill(buf[LA][:], WHITE)
		fill(buf[RA][:], WHITE)
		fill(buf[H][:], BLACK)

		p.init = true
	}
}

func (p *Arms) Reset() {

	p.init = false
}

// 2
type Legs struct {
	init bool
}

func (p *Legs) SetLEDs(buf [][]color.RGBA, tick uint32) {

	if !p.init {
		fill(buf[LB][:], WHITE)
		fill(buf[RB][:], WHITE)
		fill(buf[LK][:], BLACK)
		fill(buf[RK][:], BLACK)
		fill(buf[LA][:], BLACK)
		fill(buf[RA][:], BLACK)
		fill(buf[H][:], BLACK)

		p.init = true
	}
}

func (p *Legs) Reset() {

	p.init = false
}

// 3
type Body struct {
	init bool
}

func (p *Body) SetLEDs(buf [][]color.RGBA, tick uint32) {

	if !p.init {
		fill(buf[LB][:], BLACK)
		fill(buf[RB][:], BLACK)
		fill(buf[LK][:], WHITE)
		fill(buf[RK][:], WHITE)
		fill(buf[LA][:], BLACK)
		fill(buf[RA][:], BLACK)
		fill(buf[H][:], BLACK)

		p.init = true
	}
}

func (p *Body) Reset() {

	p.init = false
}

// 4
type Hueshift struct {
	hue  uint16
	init bool
}

func (p *Hueshift) SetLEDs(buf [][]color.RGBA, tick uint32) {

	if !p.init {
		fill(buf[H][:], BLACK)

		p.init = true
	}

	r, g, b := Hsv2Rgb(float64(p.hue), 1.0, 1.0)
	c := color.RGBA{
		r, g, b, 255,
	}

	fill(buf[LB][:], c)
	fill(buf[RB][:], c)
	fill(buf[LK][:], c)
	fill(buf[RK][:], c)
	fill(buf[LA][:], c)
	fill(buf[RA][:], c)

	p.hue = (p.hue + 1) % 360
}

func (p *Hueshift) Reset() {

	p.init = false
	p.hue = 0
}

// 5
type Switching struct {
	pos  int8
	left bool
	init bool
}

const (
	K_OFFSET = LB_COUNT
	A_OFFSET = LB_COUNT + LK_COUNT
)

func (p *Switching) SetLEDs(buf [][]color.RGBA, tick uint32) {
	if !p.init {
		fill(buf[LB][:], BLACK)
		fill(buf[RB][:], BLACK)
		fill(buf[LK][:], BLACK)
		fill(buf[RK][:], BLACK)
		fill(buf[LA][:], BLACK)
		fill(buf[RA][:], BLACK)
		fill(buf[H][:], BLACK)

		p.init = true
	}

	if tick%5 == 0 {
		if p.pos == 56 {
			p.pos = 0
			p.left = !p.left
			fill(buf[LB][:], BLACK)
			fill(buf[RB][:], BLACK)
			fill(buf[LK][:], BLACK)
			fill(buf[RK][:], BLACK)
			fill(buf[LA][:], BLACK)
			fill(buf[RA][:], BLACK)
		}

		r, g, b := Hsv2Rgb(6.0*float64(p.pos), 1.0, 1.0)
		c := color.RGBA{
			r, g, b, 255,
		}

		// leg
		if p.pos < LB_COUNT {
			if p.left {
				buf[LB][LB_COUNT-(1+p.pos)] = c
			} else {
				buf[RB][RB_COUNT-(1+p.pos)] = c
			}
		}

		// body
		if p.pos >= LB_COUNT && p.pos < LB_COUNT+LK_COUNT {
			index := p.pos - LB_COUNT
			if p.left {
				buf[LK][index] = c
			} else {
				buf[RK][index] = c
			}
		}

		// arm
		if p.pos >= LB_COUNT+LK_COUNT && p.pos < LB_COUNT+LK_COUNT+RA_COUNT {
			index := p.pos - (LB_COUNT + LK_COUNT)
			if p.left {
				buf[RA][index] = c
			} else {
				buf[LA][index] = c
			}
		}

		p.pos = p.pos + 1
	}
}

func (p *Switching) Reset() {

	p.init = false
	p.left = false
	p.pos = 0
}

// 6
type Heart struct {
	timer    uint8
	backward bool
	init     bool
}

func (p *Heart) SetLEDs(buf [][]color.RGBA, tick uint32) {

	if !p.init {
		fill(buf[LB][:], BLACK)
		fill(buf[RB][:], BLACK)
		fill(buf[LA][:], BLACK)
		fill(buf[RA][:], BLACK)

		fill(buf[LK][:9], BLACK)
		fill(buf[RK][:9], BLACK)

		fill(buf[LK][14:], BLACK)
		fill(buf[RK][14:], BLACK)

		p.init = true
	}

	hue := 360.0 - float64(p.timer)*0.75

	r, g, b := Hsv2Rgb(hue, 1.0, 1.0)
	c := color.RGBA{
		r, g, b, 255,
	}

	fill(buf[LK][9:14], c)
	fill(buf[RK][9:14], c)
	fill(buf[H][:], c)

	if p.backward {
		p.timer = p.timer - 1
		if p.timer == 0 {
			p.backward = false
		}
	} else {
		p.timer = p.timer + 1
		if p.timer == 79 {
			p.backward = true
		}
	}
}

func (p *Heart) Reset() {

	p.init = false
	p.backward = false
	p.timer = 0
}

// 7
type Blinky struct {
	init bool
}

func (p *Blinky) SetLEDs(buf [][]color.RGBA, tick uint32) {
	if !p.init {
		fill(buf[H][:], BLACK)

		p.init = true
	}

	var c color.RGBA
	if tick%2 == 0 {
		c = WHITE
	} else {
		c = BLACK
	}

	fill(buf[LB][:], c)
	fill(buf[RB][:], c)
	fill(buf[LK][:], c)
	fill(buf[RK][:], c)
	fill(buf[LA][:], c)
	fill(buf[RA][:], c)

}

func (p *Blinky) Reset() {

	p.init = false
}

// type SolidRed struct{}

// func (p *SolidRed) SetLEDs(buf []color.RGBA, tick uint32) {
// 	for i := range buf {
// 		buf[i] = RED
// 	}
// }

// type Selective struct{}

// func (p *Selective) SetLEDs(buf []color.RGBA, tick uint32) {
// 	for i := range buf {
// 		if i >= 0 && i < 2 {
// 			buf[i] = RED
// 		} else if i >= 3 && i < 5 {
// 			buf[i] = YELLOW
// 		} else if i >= 6 && i < 8 {
// 			buf[i] = GREEN
// 		} else if i >= 9 && i < 11 {
// 			buf[i] = CYAN
// 		} else if i >= 12 && i < 14 {
// 			buf[i] = BLUE
// 		} else if i >= 15 && i < 17 {
// 			buf[i] = MAGENTA
// 		} else {
// 			buf[i] = BLACK
// 		}
// 	}
// }

// type Blinky struct{}

// func (p *Blinky) SetLEDs(buf []color.RGBA, tick uint32) {
// 	var c color.RGBA
// 	if tick%2 == 0 {
// 		c = WHITE
// 	} else {
// 		c = BLACK
// 	}

// 	for i := range buf {
// 		buf[i] = c
// 	}
// }

// type HueShift struct {
// 	hue uint16
// }

// func (p *HueShift) SetLEDs(buf []color.RGBA, tick uint32) {
// 	r, g, b := Hsv2Rgb(float64(p.hue), 1.0, 1.0)
// 	c := color.RGBA{
// 		r, g, b, 255,
// 	}

// 	for i := range buf {
// 		buf[i] = c
// 	}

// 	p.hue = (p.hue + 1) % 360
// }

// type Snake struct{}

// func (p *Snake) SetLEDs(buf []color.RGBA, tick uint32) {
// 	tick = tick / 2
// 	total := uint32(len(buf)) + 4

// 	pos := tick % total

// 	cycle := (tick % (6 * total)) / total

// 	var head, body, bg color.RGBA

// 	switch cycle {
// 	case 0:
// 		head = RED
// 		body = YELLOW
// 		bg = CYAN
// 	case 1:
// 		head = YELLOW
// 		body = GREEN
// 		bg = BLUE
// 	case 2:
// 		head = GREEN
// 		body = CYAN
// 		bg = MAGENTA
// 	case 3:
// 		head = CYAN
// 		body = BLUE
// 		bg = RED
// 	case 4:
// 		head = BLUE
// 		body = MAGENTA
// 		bg = YELLOW
// 	case 5:
// 		head = MAGENTA
// 		body = RED
// 		bg = GREEN
// 	}

// 	// background
// 	for i := range buf {
// 		buf[i] = bg
// 	}

// 	// head
// 	if pos >= 2 && pos < total-2 {
// 		buf[pos-2] = head
// 	}

// 	// body a
// 	if pos >= 3 && pos < total-1 {
// 		buf[pos-3] = body
// 	}
// 	// body b
// 	if pos >= 4 && pos < total {
// 		buf[pos-4] = body
// 	}
// }
