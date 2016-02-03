package enigma

import (
	"errors"
	//"fmt"
)

type Enigma struct {
	rotors    []*Rotor
	reflector *Reflector
	// amount of pins that have every rotor and reflector
	power int
}

var (
	ErrNoRotors    = errors.New("enigma: no rotors")
	ErrNoReflector = errors.New("enigma: no reflector")

	ErrInvalidComponent = errors.New("enigma: invalid rotor/reflector, power must be same")

	ErrInValOutOfRange = errors.New("enigma: can't transform input value, out of range")
)

// Create new enigma with reflector and rotors.
// Reflector and rotors must have same amount of pins (power).
func NewEnigma(rotors []*Rotor, reflector *Reflector) (*Enigma, error) {
	if len(rotors) == 0 {
		return nil, ErrNoRotors
	}
	if reflector == nil {
		return nil, ErrNoReflector
	}

	power := reflector.Power()
	for _, r := range rotors {
		if r.power != power {
			return nil, ErrInvalidComponent
		}
	}

	return &Enigma{
		rotors:    rotors,
		reflector: reflector,
		power:     power,
	}, nil
}

// Return amount of pins that have every rotor and reflector.
func (e *Enigma) Power() int {
	return e.power
}

// Produce transfomation through rotors and reflector.
func (e *Enigma) Transform(in int) (out int, err error) {
	if in < 0 || e.power <= in {
		return 0, ErrInValOutOfRange
	}
	e.turn()
	out = e.forwardTransformation(in)
	out = e.reflector.transform(out)
	out = e.backwardTransformation(out)
	return out, nil
}

func (e *Enigma) turn() {
	//fmt.Printf("Turning...\n\tTurn rotor 0 by default\n")
	e.rotors[0].turn()
	for i := 0; i < len(e.rotors)-1; i++ {
		if e.rotors[i].isTurnoverPos() {
			//fmt.Printf("\tTurn rotor %v\n", i+1)
			e.rotors[i+1].turn()
		}
	}
}

func (e *Enigma) forwardTransformation(in int) (out int) {
	//fmt.Printf("Forward transformation value %v\n", in)
	out = in
	for _, r := range e.rotors {
		//fmt.Printf("\trotor %v: %v -> ", i, out)
		out = r.transformToLeft(out)
		//fmt.Printf("%v\n", out)
	}
	return out
}

func (e *Enigma) backwardTransformation(in int) (out int) {
	//fmt.Printf("Backward transformation value %v\n", in)
	out = in
	for i := len(e.rotors) - 1; i >= 0; i-- {
		//fmt.Printf("\trotor %v: %v -> ", i, out)
		out = e.rotors[i].transformToRight(out)
		//fmt.Printf("%v\n", out)
	}
	return out
}
