package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/IU7Abyss/cryptography/02-enigma/enigma"
)

const (
	kSeed    = 333
	kPower   = 12
	kNRotors = 3
)

var (
	ErrInvalidNumberOfRotors = errors.New("example: invalid number of rotors")
	ErrInvalidPower          = errors.New("example: invalid power of rotor")
)
var (
	kTurnPoss = [kNRotors][]int{
		[]int{2},
		[]int{2},
		[]int{10},
	}
	kPoss  = [kNRotors]int{0, 0, 0}
	kRings = [kNRotors]int{0, 0, 0}
)

func main() {
	// build random rotors and reflector
	randSrc := rand.NewSource(kSeed)
	rotors, err := genRandRotors(randSrc, kPower, kTurnPoss[:], kPoss[:], kRings[:], kNRotors)
	if err != nil {
		fmt.Println(err)
		return
	}
	reflector, err := genRandReflector(randSrc, kPower)
	if err != nil {
		fmt.Println(err)
		return
	}
	randEnigma, err := enigma.NewEnigma(rotors, reflector)
	if err != nil {
		fmt.Println(err)
		return
	}

	var out int
	//values := []int{0, 1, 2}
	//values := []int{3, 1, 8}
	//values := []int{11, 1, 10}
	values := []int{2, 9, 0}
	for _, v := range values {
		if out, err = randEnigma.Transform(v); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v, out)
	}
}

func genRandRules(src rand.Source, power int) (rules []int, err error) {
	if power <= 0 {
		return nil, ErrInvalidPower
	}
	r := rand.New(src)
	rules = r.Perm(power)
	marks := make([]bool, len(rules))
	for i, v := range rules {
		if marks[i] || marks[v] {
			continue
		}
		marks[i], marks[v] = true, true
		rules[v] = i
	}
	for i, isMarked := range marks {
		if isMarked {
			continue
		}
		rules[i] = i
	}
	fmt.Printf("rules: %#v\n", rules)
	return rules, nil
}

func genRandRotors(src rand.Source, power int, turnPoss [][]int, poss []int, rings []int, n int) (rotors []*enigma.Rotor, err error) {
	if n <= 0 {
		return nil, ErrInvalidNumberOfRotors
	}

	rotors = make([]*enigma.Rotor, n)
	for i, rotor := range rotors {
		rules, err := genRandRules(src, power)
		if err != nil {
			return nil, err
		}
		rotor, err = enigma.NewRotor(rules, poss[i], turnPoss[i], rings[i])
		if err != nil {
			return nil, err
		}
		rotors[i] = rotor
	}
	return
}

func genRandReflector(src rand.Source, power int) (reflector *enigma.Reflector, err error) {
	rules, err := genRandRules(src, power)
	if err != nil {
		return nil, err
	}
	return enigma.NewReflector(rules)
}
