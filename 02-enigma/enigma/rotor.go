package enigma

import (
	"errors"
)

func modulus(a, b int) int { return ((a % b) + b) % b }

type Rotor struct {
	// Pins
	w *wiring

	//
	pos int

	//
	turnoverPoss map[int]bool

	//
	ring int

	// Amount of pins
	power int
}

var (
	ErrNegPos         = errors.New("rotor: position must be positive")
	ErrNegRing        = errors.New("rotor: ring setting must be positive")
	ErrNegTurnoverPos = errors.New("rotor: turnover position must be positive")

	ErrNoRules        = errors.New("rotor: no rules")
	ErrNoTurnoverPoss = errors.New("rotor: no turnover positions")

	ErrInvalidPos         = errors.New("rotor: position must be less power of rotor")
	ErrInvalidRing        = errors.New("rotor: ring setting must be less the power of rotor")
	ErrInvalidTurnoverPos = errors.New("rotor: turnover position must be less power of rotor")
)

func NewRotor(rules []int, pos int, turnoverPoss []int, ring int) (r *Rotor, err error) {
	power := len(rules)
	if power == 0 {
		return nil, ErrNoRules
	}
	if len(turnoverPoss) == 0 {
		return nil, ErrNoTurnoverPoss
	}
	if pos < 0 {
		return nil, ErrNegPos
	}
	if ring < 0 {
		return nil, ErrNegRing
	}
	if pos >= power {
		return nil, ErrInvalidPos
	}
	if ring >= power {
		return nil, ErrInvalidRing
	}

	turnPoss := map[int]bool{}
	for _, turnPos := range turnoverPoss {
		if turnPos < 0 {
			return nil, ErrNegTurnoverPos
		}
		if turnPos >= power {
			return nil, ErrInvalidTurnoverPos
		}
		turnPoss[turnPos] = true
	}

	return &Rotor{
		w:            newWiring(rules),
		pos:          pos,
		turnoverPoss: turnPoss,
		ring:         ring,
		power:        power,
	}, nil
}

func (r *Rotor) SetPos(pos int) error {
	if pos < 0 {
		return ErrNegPos
	}
	if pos >= r.power {
		return ErrInvalidPos
	}

	r.pos = pos

	return nil
}

func (r *Rotor) SetRing(ring int) error {
	if ring < 0 {
		return ErrNegRing
	}
	if ring >= r.power {
		return ErrInvalidRing
	}

	r.ring = ring

	return nil
}

func (r *Rotor) turn() {
	r.pos = modulus(r.pos+1, r.power)
}

func (r *Rotor) isTurnoverPos() bool {
	return r.turnoverPoss[r.pos]
}

func (r *Rotor) Pos() int   { return r.pos }
func (r *Rotor) Ring() int  { return r.ring }
func (r *Rotor) Power() int { return r.power }

func (r *Rotor) AddTurnoverPos(pos int) error {
	if pos < 0 {
		return ErrNegTurnoverPos
	}
	if pos >= r.power {
		return ErrInvalidTurnoverPos
	}

	r.turnoverPoss[pos] = true

	return nil
}

func (r *Rotor) RemoveTurnoverPos(pos int) {
	delete(r.turnoverPoss, pos)
}

func (r *Rotor) transformToLeft(in int) (out int) {
	return r.transform(in, r.w.rightPins)
}

func (r *Rotor) transformToRight(in int) (out int) {
	return r.transform(in, r.w.leftPins)
}

func (r *Rotor) transform(in int, rules []int) (out int) {
	shift := r.pos - r.ring
	i := modulus(in+shift, r.power)
	out = modulus(rules[i]-shift, r.power)

	return out
}
