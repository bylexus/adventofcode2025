package day06

import (
	"fmt"
	"strings"

	"alexi.ch/aoc/2025/lib"
)

type Column struct {
	op       rune
	startIdx int
	endIdx   int
}

type Day06 struct {
	s1               int
	s2               int
	numbers          [][]int
	ops              []string
	prob2Cols        []Column
	prob2NumberLines []string
}

func New() Day06 {
	return Day06{s1: 0, s2: 0, numbers: make([][]int, 0), ops: make([]string, 0), prob2Cols: make([]Column, 0), prob2NumberLines: make([]string, 0)}
}

func (d *Day06) Title() string {
	return "Day 06 - TRASH COMPACTOR"
}

func (d *Day06) Setup() {
	// var lines = lib.ReadLines("data/06-test-data.txt")
	var lines = lib.ReadLines("data/06-data.txt")
	// part 1:
	for i := 0; i < len(lines)-1; i++ {
		line := lines[i]
		splits := strings.Fields(line)
		nrs := lib.Map(&splits, func(s string) int { return lib.StrToInt(s) })
		d.numbers = append(d.numbers, nrs)
		d.prob2NumberLines = append(d.prob2NumberLines, line)
	}
	d.ops = strings.Fields(lines[len(lines)-1])

	// part 2:
	// we figure out the start/end index for each column,
	// and just remember those indices.

	opsLine := lines[len(lines)-1]
	width := len(opsLine)
	endIdx := width - 1
	for i := width - 1; i >= 0; i-- {
		if c := rune(opsLine[i]); c != ' ' {
			op := c
			startIdx := i
			col := Column{
				op:       op,
				startIdx: startIdx,
				endIdx:   endIdx,
			}
			d.prob2Cols = append(d.prob2Cols, col)
			endIdx = startIdx - 2
		}
	}
}

func (d *Day06) SolveProblem1() {
	d.s1 = 0
	for col := 0; col < len(d.ops); col++ {
		total := 0
		if d.ops[col] == "*" {
			total = 1
		}
		for row := 0; row < len(d.numbers); row++ {
			switch d.ops[col] {
			case "+":
				total += d.numbers[row][col]
			case "*":
				total *= d.numbers[row][col]
			}
		}
		d.s1 += total
	}
}

func (d *Day06) SolveProblem2() {
	d.s2 = 0
	height := len(d.prob2NumberLines)
	for _, col := range d.prob2Cols {
		total := 0
		if col.op == '*' {
			total = 1
		}
		// nrIdx is the column index of the actual number:
		for nrIdx := col.endIdx; nrIdx >= col.startIdx; nrIdx-- {
			number := 0
			// append number in one column:
			for line := 0; line < height; line++ {
				if r := rune(d.prob2NumberLines[line][nrIdx]); r != ' ' {
					number = (number * 10) + lib.StrToInt(string(r))
				}
			}

			// sum/multiply:
			switch col.op {
			case '+':
				total += number
			case '*':
				total *= number
			}
		}
		d.s2 += total
	}
}

func (d *Day06) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day06) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
