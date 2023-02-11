package led

import (
	"image/color"

	"machine"

	"tinygo.org/x/drivers/ws2812"
)

const (
	leftLegPin   = machine.GP14
	leftBodyPin  = machine.GP12
	leftArmPin   = machine.GP10
	rightLegPin  = machine.GP15
	rightBodyPin = machine.GP13
	rightArmPin  = machine.GP11
	heartPin     = machine.GP9
)

type LedSuit struct {
	strips  []ws2812.Device
	buffers [][]color.RGBA
	count   *LedCount
}

func NewLedSuit(pid string) *LedSuit {
	count := GetLedCount(pid)

	strips := make([]ws2812.Device, stripCount)

	leftLegPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[LeftLeg] = ws2812.New(leftLegPin)
	leftBodyPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[LeftBody] = ws2812.New(leftBodyPin)
	leftArmPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[LeftArm] = ws2812.New(leftArmPin)
	rightLegPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[RightLeg] = ws2812.New(rightLegPin)
	rightBodyPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[RightBody] = ws2812.New(rightBodyPin)
	rightArmPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[RightArm] = ws2812.New(rightArmPin)
	heartPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[Heart] = ws2812.New(heartPin)

	buffers := make([][]color.RGBA, stripCount)

	buffers[LeftLeg] = make([]color.RGBA, count.Leg)
	buffers[LeftBody] = make([]color.RGBA, count.Body)
	buffers[LeftArm] = make([]color.RGBA, count.Arm)
	buffers[RightLeg] = make([]color.RGBA, count.Leg)
	buffers[RightBody] = make([]color.RGBA, count.Body)
	buffers[RightArm] = make([]color.RGBA, count.Arm)
	buffers[Heart] = make([]color.RGBA, count.Heart)

	return &LedSuit{
		strips:  strips,
		buffers: buffers,
		count:   count,
	}
}

func (l *LedSuit) GetBuffer(lid LID) []color.RGBA {
	return l.buffers[lid]
}

func (l *LedSuit) GetCount() *LedCount {
	return l.count
}

func (l *LedSuit) Display() {
	for i := 0; i < int(stripCount); i++ {
		l.strips[i].WriteColors(l.buffers[i])
	}
}
