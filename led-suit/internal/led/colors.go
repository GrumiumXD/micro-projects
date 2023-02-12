package led

import (
	"image/color"
	"math"
)

var BRIGHTNESS float64 = 0.25

var RED = color.RGBA{
	uint8(255 * BRIGHTNESS),
	0,
	0,
	255,
}
var GREEN = color.RGBA{
	0,
	uint8(255 * BRIGHTNESS),
	0,
	255,
}
var BLUE = color.RGBA{
	0,
	0,
	uint8(255 * BRIGHTNESS),
	255,
}
var MAGENTA = color.RGBA{
	uint8(255 * BRIGHTNESS),
	0,
	uint8(255 * BRIGHTNESS),
	255,
}
var YELLOW = color.RGBA{
	uint8(255 * BRIGHTNESS),
	uint8(255 * BRIGHTNESS),
	0,
	255,
}
var ORANGE = color.RGBA{
	uint8(255 * BRIGHTNESS),
	uint8(128 * BRIGHTNESS),
	0,
	255,
}
var BRIGHT_GREEN = color.RGBA{
	uint8(128 * BRIGHTNESS),
	uint8(255 * BRIGHTNESS),
	0,
	255,
}
var CYAN = color.RGBA{
	0,
	uint8(255 * BRIGHTNESS),
	uint8(255 * BRIGHTNESS),
	255,
}
var WHITE = color.RGBA{
	uint8(255 * BRIGHTNESS),
	uint8(255 * BRIGHTNESS),
	uint8(255 * BRIGHTNESS),
	255,
}
var BLACK = color.RGBA{
	0,
	0,
	0,
	255,
}

func FillBuffer(buf []color.RGBA, color color.RGBA) {
	buf[0] = color

	for i := 1; i < len(buf); i *= 2 {
		copy(buf[i:], buf[:i])
	}
}

func Hsv2Rgb(hue float64, sat float64, val float64) (uint8, uint8, uint8) {
	c := val * sat
	x := c * (1.0 - math.Abs(math.Mod(hue/60.0, 2.0)-1.0))
	m := val - c

	var r, g, b float64

	if hue >= 0.0 && hue < 60.0 {
		r = c
		g = x
		b = 0
	} else if hue >= 60.0 && hue < 120.0 {
		r = x
		g = c
		b = 0
	} else if hue >= 120.0 && hue < 180.0 {
		r = 0
		g = c
		b = x
	} else if hue >= 180.0 && hue < 240.0 {
		r = 0
		g = x
		b = c
	} else if hue >= 240.0 && hue < 300.0 {
		r = x
		g = 0
		b = c
	} else if hue >= 300.0 {
		r = c
		g = 0
		b = x
	}

	return uint8((r + m) * BRIGHTNESS * 255), uint8((g + m) * BRIGHTNESS * 255), uint8((b + m) * BRIGHTNESS * 255)
}
