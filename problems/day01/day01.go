package day01

import (
	"fmt"
	"regexp"

	"alexi.ch/aoc/2025/lib"
)

type Input struct {
	Dir    int
	Amount int
}

type Day01 struct {
	s1    int
	s2    int
	input []Input
}

func New() Day01 {
	return Day01{input: make([]Input, 0)}
}

func (d *Day01) Title() string {
	return "Day 01 - Secret Entrance"
}

func (d *Day01) Setup() {
	var lines = lib.ReadLines("data/01-data.txt")
	// var lines = lib.ReadLines("data/01-test-data.txt")
	for _, line := range lines {
		r := regexp.MustCompile(`(L|R)(\d+)`)
		matches := r.FindStringSubmatch(line)
		// fmt.Printf("line: %#v\n", matches)
		if len(matches) == 3 {
			input := Input{}
			if matches[1] == "L" {
				input.Dir = -1
			} else if matches[1] == "R" {
				input.Dir = 1
			} else {
				panic("Unknown dir")
			}
			input.Amount = lib.StrToInt(matches[2])
			d.input = append(d.input, input)
		}
	}
}

func (d *Day01) SolveProblem1() {
	dial := 50
	zeroes := 0
	for _, input := range d.input {
		// fmt.Printf("Before: %d, ", dial)
		amount := input.Amount * input.Dir
		// fmt.Printf("amount: %d, ", amount)
		newDial := (dial + amount) % 100
		if newDial < 0 {
			dial = 100 + newDial
		} else {
			dial = newDial
		}
		// fmt.Printf("after: %d\n", dial)
		if dial == 0 {
			zeroes++
		}
	}
	d.s1 = zeroes
}

func (d *Day01) SolveProblem2() {
	// I'm sure there would be a more elegant solution ....
	// bloody modulus and off-by-one's :-)

	dial := 50
	zeroes := 0
	for _, input := range d.input {
		amount := input.Amount * input.Dir
		rest := (dial + amount)
		zeroPassings := 0
		// we dial backwards over 0:
		if rest < 0 {
			zeroPassings = lib.Abs((dial + amount) / 100)
			// if we were on 0 already, we don't count the 1st zero passing. Else, we
			// have to count the 1st zero passing, too:
			if dial > 0 {
				zeroPassings++
			}
		} else if rest == 0 {
			// we landed on zero:
			zeroPassings++
		} else {
			// we dialed to the right:
			zeroPassings = lib.Abs((dial + amount) / 100)
		}

		zeroes += zeroPassings
		rest = rest % 100
		if rest < 0 {
			dial = 100 + rest
		} else {
			dial = rest
		}
	}
	d.s2 = zeroes
}

func (d *Day01) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day01) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
