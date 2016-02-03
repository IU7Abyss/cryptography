package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/IU7Abyss/cryptography/02-enigma/enigma"
)

var (
	ErrNoFile = errors.New("example: rand-enigma inputfile outputfile")

	ErrInvalidNumberOfRotors = errors.New("example: invalid number of rotors")
	ErrInvalidPower          = errors.New("example: invalid power of rotor")
	ErrInvalidFile           = errors.New("example: invalid input file")
)

const (
	kSeed    = 333
	kPower   = 256
	kNRotors = 3
)

var (
	kTurnPoss = [kNRotors][]int{
		[]int{5},
		[]int{2},
		[]int{10},
	}
	kPoss  = [kNRotors]int{40, 10, 25}
	kRings = [kNRotors]int{25, 30, 10}
)

func main() {
	// path to file should be
	if len(os.Args) < 3 {
		fmt.Println(ErrNoFile)
		return
	}
	inFilename := os.Args[1]
	outFilename := os.Args[2]

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

	// transform input file
	inFile, err := ioutil.ReadFile(inFilename)
	if err != nil {
		fmt.Println(ErrInvalidFile)
		return
	}
	var out int
	outFile := make([]byte, len(inFile))
	for i, in := range inFile {
		if out, err = randEnigma.Transform(int(in)); err != nil {
			fmt.Println(err)
			return
		}
		outFile[i] = byte(out)
	}
	if err = ioutil.WriteFile(outFilename, outFile, 0644); err != nil {
		fmt.Println(err)
		return
	}
	return
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
