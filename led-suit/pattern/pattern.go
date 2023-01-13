package pattern

type Pattern interface {
	SetLEDs(tick uint32)
	Reset()
}
