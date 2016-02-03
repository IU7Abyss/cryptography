package enigma

import (
	"errors"
)

type Reflector struct {
	rules []int
}

var (
	ErrReflNoRules = errors.New("reflector: no rules")
)

func NewReflector(rules []int) (*Reflector, error) {
	if len(rules) == 0 {
		return nil, ErrReflNoRules
	}
	return &Reflector{
		rules: rules,
	}, nil
}

func (r *Reflector) transform(in int) (out int) {
	return r.rules[in]
}

func (r *Reflector) Power() int {
	return len(r.rules)
}
