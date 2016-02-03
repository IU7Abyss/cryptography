package enigma

type wiring struct {
	leftPins  []int
	rightPins []int
}

func newWiring(rightPins []int) *wiring {
	leftPins := make([]int, len(rightPins))
	for pin, i := range rightPins {
		leftPins[i] = pin
	}
	return &wiring{
		leftPins:  leftPins,
		rightPins: rightPins,
	}
}
