package led

type LID int

const (
	LeftLeg LID = iota
	LeftBody
	LeftArm
	RightLeg
	RightBody
	RightArm
	Heart

	stripCount
)

type LedCount struct {
	Leg   int8
	Body  int8
	Arm   int8
	Heart int8
	// first heart LED on the body strip
	HBStart int8
	// last heart LED on the body strip
	HBEnd int8
}

func GetLedCount(pid string) *LedCount {
	switch pid {
	case "GR":
		return &LedCount{
			Leg:     22,
			Body:    19,
			Arm:     15,
			Heart:   9,
			HBStart: 9,
			HBEnd:   13,
		}
	case "RS":
		return &LedCount{
			Leg:     25,
			Body:    22,
			Arm:     16,
			Heart:   9,
			HBStart: 9,
			HBEnd:   13,
		}
	case "BK":
		return &LedCount{
			Leg:     24,
			Body:    21,
			Arm:     16,
			Heart:   9,
			HBStart: 9,
			HBEnd:   13,
		}
	case "GB":
		return &LedCount{
			Leg:     24,
			Body:    21,
			Arm:     16,
			Heart:   9,
			HBStart: 9,
			HBEnd:   13,
		}
	}

	// default
	return &LedCount{}
}
