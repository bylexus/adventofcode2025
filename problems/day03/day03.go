package day03

import (
	"fmt"
	"strconv"

	"alexi.ch/aoc/2025/lib"
)

type Day03 struct {
	s1    int
	s2    int
	banks [][]int
}

func New() Day03 {
	return Day03{s1: 0, s2: 0, banks: make([][]int, 0)}
}

func (d *Day03) Title() string {
	return "Day 03 - Lobby"
}

func (d *Day03) Setup() {
	// var lines = lib.ReadLines("data/03-test-data.txt")
	var lines = lib.ReadLines("data/03-data.txt")
	for _, line := range lines {
		bank := make([]int, 0)
		for _, char := range line {
			nr, err := strconv.Atoi(string(char))
			lib.Check(err)
			bank = append(bank, nr)
		}
		d.banks = append(d.banks, bank)
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day03) SolveProblem1() {
	d.s1 = d.solveForNrOfBatteries(2)
}

func (d *Day03) SolveProblem2() {
	d.s2 = d.solveForNrOfBatteries(12)
}

// Idea: (works for both solutions:)
// we need to find the most-left highest possible value for each of the n digit positions:
//   - The first digit must be within position 0 to (len-n): It cannot be farther right than len-n.
//     So we look for the highest leftmost number in this range.
//   - The 2nd digit can only be right of the first one, but (len-n+1) to the right.
//     Also here, we search the leftmost highest number in that range.
//   - Same for 3rd, 4th, and so on.
func (d *Day03) solveForNrOfBatteries(nrOfBatteries int) int {
	sum := 0
	for _, bank := range d.banks {
		minPos := 0
		maxPos := len(bank) - nrOfBatteries
		maxBankValue := 0
		for actDigit := 0; actDigit < nrOfBatteries; actDigit++ {
			maxValue := 0
			// search for biggest, leftmost digit within allowed range:
			for idx := minPos; idx <= maxPos; idx++ {
				if bank[idx] > maxValue {
					minPos = idx
					maxValue = bank[idx]
				}
			}
			maxBankValue = 10*maxBankValue + maxValue

			// prepare for next digit:
			minPos = minPos + 1
			maxPos = maxPos + 1
		}
		sum += maxBankValue
	}
	return sum
}

func (d *Day03) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day03) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
