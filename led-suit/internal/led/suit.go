package led

import (
	"image/color"
)

type LedSuit interface {
	GetBuffer(lid LID) []color.RGBA
	GetCount() *LedCount
	Display()
}
