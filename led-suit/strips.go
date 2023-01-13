package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ws2812"
)

const (
	LB = iota
	LK
	LA
	RB
	RK
	RA
	H
	STRIP_COUNT
)

const (
	LB_COUNT = 22
	LK_COUNT = 19
	LA_COUNT = 15
	RB_COUNT = LB_COUNT
	RK_COUNT = LK_COUNT
	RA_COUNT = LA_COUNT
	H_COUNT  = 9
)

const (
	LB_PIN = machine.GP14
	RB_PIN = machine.GP15
	LK_PIN = machine.GP12
	RK_PIN = machine.GP13
	LA_PIN = machine.GP10
	RA_PIN = machine.GP11
	H_PIN  = machine.GP9
)

func stripSetup() ([]ws2812.Device, [][]color.RGBA) {
	strips := make([]ws2812.Device, STRIP_COUNT)

	LB_PIN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[LB] = ws2812.New(LB_PIN)
	LK_PIN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[LK] = ws2812.New(LK_PIN)
	LA_PIN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[LA] = ws2812.New(LA_PIN)
	RB_PIN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[RB] = ws2812.New(RB_PIN)
	RK_PIN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[RK] = ws2812.New(RK_PIN)
	RA_PIN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[RA] = ws2812.New(RA_PIN)
	H_PIN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	strips[H] = ws2812.New(H_PIN)

	buffers := make([][]color.RGBA, STRIP_COUNT)

	buffers[LB] = make([]color.RGBA, LB_COUNT)
	buffers[LK] = make([]color.RGBA, LK_COUNT)
	buffers[LA] = make([]color.RGBA, LA_COUNT)
	buffers[RB] = make([]color.RGBA, RB_COUNT)
	buffers[RK] = make([]color.RGBA, RK_COUNT)
	buffers[RA] = make([]color.RGBA, RA_COUNT)
	buffers[H] = make([]color.RGBA, H_COUNT)

	return strips, buffers
}
