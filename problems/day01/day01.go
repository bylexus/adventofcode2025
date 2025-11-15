package day01

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"alexi.ch/aoc/2025/lib"
	"github.com/bylexus/go-stdlib/eerr"
)

type Day01 struct {
	leftNumbers  []int
	rightNumbers []int

	s1 int
	s2 int
}

func New() Day01 {
	return Day01{leftNumbers: make([]int, 0), rightNumbers: make([]int, 0)}
}

func (d *Day01) Title() string {
	return "Day 01 - Historian Hysteria"
}

func (d *Day01) Setup() {
	var lines = lib.ReadLines("data/01-data.txt")
	// var lines = lib.ReadLines("data/01-test-data.txt")
	for _, line := range lines {
		r := regexp.MustCompile(`(\d+)\s+(\d+)`)
		matches := r.FindStringSubmatch(line)
		// fmt.Printf("line: %#v\n", matches)
		if len(matches) == 3 {
			leftNr, err := strconv.ParseUint(matches[1], 10, 64)
			eerr.PanicOnErr(err)
			rightNumber, err := strconv.ParseUint(matches[2], 10, 64)
			eerr.PanicOnErr(err)
			d.leftNumbers = append(d.leftNumbers, int(leftNr))
			d.rightNumbers = append(d.rightNumbers, int(rightNumber))
		}
	}
	// fmt.Printf("%#v\n", d.leftNumbers)
	// fmt.Printf("%#v\n", d.rightNumber)
}

func (d *Day01) SolveProblem1() {
	// Idea:
	// first, sort both lists.
	// Then just walk the left list, and compare the number on the right,
	// calc distances and sum up:

	// sort both input arrays
	slices.Sort(d.leftNumbers)
	slices.Sort(d.rightNumbers)
	sum := 0
	for i := range d.leftNumbers {
		d := lib.Abs(d.leftNumbers[i] - d.rightNumbers[i])
		sum += d
	}
	d.s1 = sum
}

func (d *Day01) SolveProblem2() {
	// Idea:
	// walk through the left list, and count the entries on the right with the same number.
	// Cache the result in a map, to avoid double counting.
	// Just prod / sum them up in the process:
	nrMap := make(map[int]int)
	sum := 0
	for i := range d.leftNumbers {
		leftNr := d.leftNumbers[i]
		countRight, present := nrMap[d.leftNumbers[i]]
		if !present {
			countRight = countNrs(d.rightNumbers, leftNr)
			nrMap[leftNr] = countRight
		}
		sum += leftNr * countRight
	}
	d.s2 = sum
}

func countNrs(list []int, n int) int {
	count := 0
	for _, entry := range list {
		if entry == n {
			count++
		}
	}
	return count
}

func (d *Day01) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day01) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
