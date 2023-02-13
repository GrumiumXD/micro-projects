package pattern

import "ledsuit/internal/led"

type Kit struct {
	suit led.LedSuit
	init bool
	pos  int
	out  bool
}

const kLength = 5

func NewKitPattern(suit led.LedSuit) Pattern {
	return &Kit{
		suit: suit,
		init: false,
		out:  true,
		pos:  -kLength,
	}
}

func (k *Kit) SetLEDs(tick uint32) {

	if !k.init {
		led.FillBuffer(k.suit.GetBuffer(led.LeftLeg)[:], led.BLACK)
		led.FillBuffer(k.suit.GetBuffer(led.LeftBody)[:], led.BLACK)
		led.FillBuffer(k.suit.GetBuffer(led.LeftArm)[:], led.BLACK)
		led.FillBuffer(k.suit.GetBuffer(led.RightLeg)[:], led.BLACK)
		led.FillBuffer(k.suit.GetBuffer(led.RightBody)[:], led.BLACK)
		led.FillBuffer(k.suit.GetBuffer(led.RightArm)[:], led.BLACK)
		led.FillBuffer(k.suit.GetBuffer(led.Heart)[:], led.BLACK)

		k.init = true
	}

	if k.pos < -kLength {
		return
	}

	// arms
	for i := 0; i < int(k.suit.GetCount().Arm); i++ {
		if i >= k.pos && i < k.pos+kLength {
			k.suit.GetBuffer(led.LeftArm)[i] = led.RED
			k.suit.GetBuffer(led.RightArm)[i] = led.RED
		} else {
			k.suit.GetBuffer(led.LeftArm)[i] = led.BLACK
			k.suit.GetBuffer(led.RightArm)[i] = led.BLACK
		}
	}

	if k.out {
		k.pos++
		if k.pos == int(k.suit.GetCount().Arm)-1 {
			k.out = false
		}
	} else {
		k.pos--
	}

}

func (k *Kit) Reset() {

	k.init = false
	k.out = true
	k.pos = -kLength
}
