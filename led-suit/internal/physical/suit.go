package physical

import (
	"image/color"
	"ledsuit/internal/led"

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
	count   *led.LedCount
}

func NewLedSuit(pid string) led.LedSuit {
	count := led.GetLedCount(pid)

	strips := make([]ws2812.Device, led.StripCount)

	leftLegPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[led.LeftLeg] = ws2812.New(leftLegPin)
	leftBodyPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[led.LeftBody] = ws2812.New(leftBodyPin)
	leftArmPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[led.LeftArm] = ws2812.New(leftArmPin)
	rightLegPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[led.RightLeg] = ws2812.New(rightLegPin)
	rightBodyPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[led.RightBody] = ws2812.New(rightBodyPin)
	rightArmPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[led.RightArm] = ws2812.New(rightArmPin)
	heartPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[led.Heart] = ws2812.New(heartPin)

	buffers := make([][]color.RGBA, led.StripCount)

	buffers[led.LeftLeg] = make([]color.RGBA, count.Leg)
	buffers[led.LeftBody] = make([]color.RGBA, count.Body)
	buffers[led.LeftArm] = make([]color.RGBA, count.Arm)
	buffers[led.RightLeg] = make([]color.RGBA, count.Leg)
	buffers[led.RightBody] = make([]color.RGBA, count.Body)
	buffers[led.RightArm] = make([]color.RGBA, count.Arm)
	buffers[led.Heart] = make([]color.RGBA, count.Heart)

	return &LedSuit{
		strips:  strips,
		buffers: buffers,
		count:   count,
	}
}

func (l *LedSuit) GetBuffer(lid led.LID) []color.RGBA {
	return l.buffers[lid]
}

func (l *LedSuit) GetCount() *led.LedCount {
	return l.count
}

func (l *LedSuit) Display() {
	for i := 0; i < int(led.StripCount); i++ {
		l.strips[i].WriteColors(l.buffers[i])
	}
}
