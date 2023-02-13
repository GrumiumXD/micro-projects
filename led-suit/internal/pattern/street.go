package pattern

import "ledsuit/internal/led"

type Street struct {
	suit led.LedSuit
	init bool
	pos  int
	tick uint32
}

const sLength = 4
const sDistance = 7
const sCycle = sLength + sDistance

func NewStreetPattern(suit led.LedSuit) Pattern {
	return &Street{
		suit: suit,
		init: false,
		pos:  0,
		tick: 0,
	}
}

func (s *Street) SetLEDs(tick uint32) {

	if !s.init {
		led.FillBuffer(s.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
		led.FillBuffer(s.suit.GetBuffer(led.LeftBody)[:], led.BLACK)
		led.FillBuffer(s.suit.GetBuffer(led.LeftArm)[:], led.BLACK)
		led.FillBuffer(s.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
		led.FillBuffer(s.suit.GetBuffer(led.RightBody)[:], led.BLACK)
		led.FillBuffer(s.suit.GetBuffer(led.RightArm)[:], led.BLACK)
		led.FillBuffer(s.suit.GetBuffer(led.Heart)[:], led.BLACK)

		s.init = true
	}

	if s.tick%4 == 0 {

		shift := sCycle - s.pos

		// arms
		for i := 0; i < int(s.suit.GetCount().Arm); i++ {
			if (i+shift)%sCycle < sLength {
				s.suit.GetBuffer(led.LeftArm)[i] = led.WHITE
				s.suit.GetBuffer(led.RightArm)[i] = led.WHITE
			} else {
				s.suit.GetBuffer(led.LeftArm)[i] = led.BLACK
				s.suit.GetBuffer(led.RightArm)[i] = led.BLACK
			}
		}

		// legs
		for i := 0; i < int(s.suit.GetCount().Leg); i++ {
			if (i+shift+int(s.suit.GetCount().Arm))%sCycle < sLength {
				s.suit.GetBuffer(led.LeftLeg)[i] = led.WHITE
				s.suit.GetBuffer(led.RightLeg)[i] = led.WHITE
			} else {
				s.suit.GetBuffer(led.LeftLeg)[i] = led.BLACK
				s.suit.GetBuffer(led.RightLeg)[i] = led.BLACK
			}
		}

		s.pos = (s.pos + 1) % sCycle
	}

	s.tick++
}

func (s *Street) Reset() {

	s.init = false
	s.pos = 0
	s.tick = 0
}
