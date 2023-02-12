package main

import (
	"image/color"
	"ledsuit/internal/led"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Suit struct {
	strips  [][]*ebiten.Image
	buffers [][]color.RGBA
	count   *led.LedCount
}

func NewSuit(pid string) *Suit {
	count := led.GetLedCount(pid)

	strips := make([][]*ebiten.Image, led.StripCount)

	strips[led.LeftLeg] = make([]*ebiten.Image, count.Leg)
	strips[led.LeftBody] = make([]*ebiten.Image, count.Body)
	strips[led.LeftArm] = make([]*ebiten.Image, count.Arm)
	strips[led.RightLeg] = make([]*ebiten.Image, count.Leg)
	strips[led.RightBody] = make([]*ebiten.Image, count.Body)
	strips[led.RightArm] = make([]*ebiten.Image, count.Arm)
	strips[led.Heart] = make([]*ebiten.Image, count.Heart)

	for i := 0; i < int(led.StripCount); i++ {
		for k := 0; k < len(strips[i]); k++ {
			strips[i][k] = ebiten.NewImage(5, 5)
		}
	}

	buffers := make([][]color.RGBA, led.StripCount)

	buffers[led.LeftLeg] = make([]color.RGBA, count.Leg)
	buffers[led.LeftBody] = make([]color.RGBA, count.Body)
	buffers[led.LeftArm] = make([]color.RGBA, count.Arm)
	buffers[led.RightLeg] = make([]color.RGBA, count.Leg)
	buffers[led.RightBody] = make([]color.RGBA, count.Body)
	buffers[led.RightArm] = make([]color.RGBA, count.Arm)
	buffers[led.Heart] = make([]color.RGBA, count.Heart)

	return &Suit{
		strips:  strips,
		buffers: buffers,
		count:   count,
	}
}

func (s *Suit) GetBuffer(lid led.LID) []color.RGBA {
	return s.buffers[lid]
}

func (s *Suit) GetCount() *led.LedCount {
	return s.count
}

func (s *Suit) Display() {
	for i := 0; i < int(led.StripCount); i++ {
		for k := 0; k < len(s.strips[i]); k++ {
			c := s.buffers[i][k]

			r := math.Min(float64(c.R)*1.0/led.BRIGHTNESS, 255.0)
			g := math.Min(float64(c.G)*1.0/led.BRIGHTNESS, 255.0)
			b := math.Min(float64(c.B)*1.0/led.BRIGHTNESS, 255.0)

			s.strips[i][k].Fill(color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
	}
}

func (s *Suit) DrawStrip(lid led.LID, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	if lid == led.LeftLeg {
		op.GeoM.Translate(450, 390)

		for i := 0; i < len(s.strips[lid]); i++ {
			op.GeoM.Translate(0, 10)

			screen.DrawImage(s.strips[lid][i], op)
		}
	} else if lid == led.RightLeg {
		op.GeoM.Translate(350, 390)

		for i := 0; i < len(s.strips[lid]); i++ {
			op.GeoM.Translate(0, 10)

			screen.DrawImage(s.strips[lid][i], op)
		}
	} else if lid == led.LeftBody {
		op.GeoM.Translate(450, 390)

		for i := 0; i < len(s.strips[lid]); i++ {
			op.GeoM.Translate(-5, -10)

			screen.DrawImage(s.strips[lid][i], op)
		}
	} else if lid == led.RightBody {
		op.GeoM.Translate(350, 390)

		for i := 0; i < len(s.strips[lid]); i++ {
			op.GeoM.Translate(5, -10)

			screen.DrawImage(s.strips[lid][i], op)
		}
	} else if lid == led.LeftArm {
		op.GeoM.Translate(470, 190)

		for i := 0; i < len(s.strips[lid]); i++ {
			op.GeoM.Translate(0, 10)

			screen.DrawImage(s.strips[lid][i], op)
		}
	} else if lid == led.RightArm {
		op.GeoM.Translate(330, 190)

		for i := 0; i < len(s.strips[lid]); i++ {
			op.GeoM.Translate(0, 10)

			screen.DrawImage(s.strips[lid][i], op)
		}
	} else if lid == led.Heart {
		op.GeoM.Translate(375, 245)

		for i := 0; i < len(s.strips[lid]); i++ {
			if i < 5 {
				op.GeoM.Translate(5, -10)
			} else {
				op.GeoM.Translate(5, 10)
			}

			screen.DrawImage(s.strips[lid][i], op)
		}
	}

}

func (s *Suit) Draw(screen *ebiten.Image) {
	s.DrawStrip(led.LeftLeg, screen)
	s.DrawStrip(led.LeftBody, screen)
	s.DrawStrip(led.LeftArm, screen)
	s.DrawStrip(led.RightLeg, screen)
	s.DrawStrip(led.RightBody, screen)
	s.DrawStrip(led.RightArm, screen)
	s.DrawStrip(led.Heart, screen)
}
