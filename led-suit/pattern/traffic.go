package pattern

import "ledsuit/led"

const (
	GREEN  = 0
	YELLOW = 1
	RED    = 2
)

type Traffic struct {
	suit  *led.LedSuit
	state int8
	init  bool
}

func NewTrafficPattern(suit *led.LedSuit) Pattern {
	return &Traffic{
		suit:  suit,
		state: 0,
		init:  false,
	}
}

func (t *Traffic) SetLEDs(tick uint32) {

	if !t.init {
		led.FillBuffer(t.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
		led.FillBuffer(t.suit.GetBuffer(led.LeftBody)[:], led.BLACK)
		led.FillBuffer(t.suit.GetBuffer(led.LeftArm)[:], led.BLACK)
		led.FillBuffer(t.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
		led.FillBuffer(t.suit.GetBuffer(led.RightBody)[:], led.BLACK)
		led.FillBuffer(t.suit.GetBuffer(led.RightArm)[:], led.BLACK)
		led.FillBuffer(t.suit.GetBuffer(led.Heart)[:], led.BLACK)

		t.init = true
	}

	count := t.suit.GetCount()
	var i int8
	// update every 20 ticks
	if tick%20 == 0 {
		led.FillBuffer(t.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
		led.FillBuffer(t.suit.GetBuffer(led.LeftBody)[:], led.BLACK)
		led.FillBuffer(t.suit.GetBuffer(led.LeftArm)[:], led.BLACK)
		led.FillBuffer(t.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
		led.FillBuffer(t.suit.GetBuffer(led.RightBody)[:], led.BLACK)
		led.FillBuffer(t.suit.GetBuffer(led.RightArm)[:], led.BLACK)

		if t.state == GREEN {
			// legs
			for i = 0; i < count.Leg; i++ {
				if i%3 == 2 {
					t.suit.GetBuffer(led.LeftLeg)[i] = led.GREEN
					t.suit.GetBuffer(led.RightLeg)[i] = led.GREEN
				}
			}

			// body
			for i = 0; i < count.Body; i++ {
				if i%3 == 0 {
					t.suit.GetBuffer(led.LeftBody)[i] = led.GREEN
					t.suit.GetBuffer(led.RightBody)[i] = led.GREEN
				}
			}

			// arms
			for i = 0; i < count.Arm; i++ {
				if i%3 == 2 {
					t.suit.GetBuffer(led.LeftArm)[i] = led.GREEN
					t.suit.GetBuffer(led.RightArm)[i] = led.GREEN
				}
			}
		}

		if t.state&YELLOW > 0 {
			// legs
			for i = 0; i < count.Leg; i++ {
				if i%3 == 1 {
					t.suit.GetBuffer(led.LeftLeg)[i] = led.YELLOW
					t.suit.GetBuffer(led.RightLeg)[i] = led.YELLOW
				}
			}

			// body
			for i = 0; i < count.Body; i++ {
				if i%3 == 1 {
					t.suit.GetBuffer(led.LeftBody)[i] = led.YELLOW
					t.suit.GetBuffer(led.RightBody)[i] = led.YELLOW
				}
			}

			// arms
			for i = 0; i < count.Arm; i++ {
				if i%3 == 1 {
					t.suit.GetBuffer(led.LeftArm)[i] = led.YELLOW
					t.suit.GetBuffer(led.RightArm)[i] = led.YELLOW
				}
			}
		}

		if t.state&RED > 0 {
			// legs
			for i = 0; i < count.Leg; i++ {
				if i%3 == 0 {
					t.suit.GetBuffer(led.LeftLeg)[i] = led.RED
					t.suit.GetBuffer(led.RightLeg)[i] = led.RED
				}
			}

			// body
			for i = 0; i < count.Body; i++ {
				if i%3 == 2 {
					t.suit.GetBuffer(led.LeftBody)[i] = led.RED
					t.suit.GetBuffer(led.RightBody)[i] = led.RED
				}
			}

			// arms
			for i = 0; i < count.Arm; i++ {
				if i%3 == 0 {
					t.suit.GetBuffer(led.LeftArm)[i] = led.RED
					t.suit.GetBuffer(led.RightArm)[i] = led.RED
				}
			}
		}

		t.state += 1
		if t.state == 4 {
			t.state = 0
		}
	}
}

func (t *Traffic) Reset() {

	t.state = 0
	t.init = false
}
