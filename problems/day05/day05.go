package day05

import (
	"fmt"
	"regexp"

	"alexi.ch/aoc/2025/lib"
)

type Day05 struct {
	s1          int
	s2          int
	freshRanges [][]int
	ingredients []int
}

func New() Day05 {
	return Day05{s1: 0, s2: 0, freshRanges: make([][]int, 0), ingredients: make([]int, 0)}
}

func (d *Day05) Title() string {
	return "Day 05 - CAFETERIA"
}

func (d *Day05) Setup() {
	// var lines = lib.ReadLines("data/05-test-data.txt")
	var lines = lib.ReadLines("data/05-data.txt")
	freshMatcher := regexp.MustCompile(`(\d+)-(\d+)`)
	ingredientMatcher := regexp.MustCompile(`^(\d+)$`)
	for _, line := range lines {
		freshRange := freshMatcher.FindStringSubmatch(line)
		ingredient := ingredientMatcher.FindStringSubmatch(line)
		if freshRange != nil {
			d.freshRanges = append(d.freshRanges, []int{
				lib.StrToInt(freshRange[1]),
				lib.StrToInt(freshRange[2]),
			})
		} else if ingredient != nil {
			d.ingredients = append(d.ingredients, lib.StrToInt(ingredient[1]))
		}
	}
}

func (d *Day05) SolveProblem1() {
	d.s1 = 0
	for _, ingredient := range d.ingredients {
		for _, freshRange := range d.freshRanges {
			if ingredient >= freshRange[0] && ingredient <= freshRange[1] {
				d.s1++
				break
			}
		}
	}
}

func (d *Day05) SolveProblem2() {
	d.s2 = 0
	// we need to calculate non-overlapping ranges from the available ranges.
	nonOverlaps := make([][]int, 0)
	// add the first range to the new array, as this one is our starting point:

	for i := 0; i < len(d.freshRanges); i++ {
		testRange := d.freshRanges[i]
		skip := false
		for j, masterRange := range nonOverlaps {
			// skip invalidated masters:
			if masterRange[0] == -1 {
				continue
			}
			// case: testRange fits into masterRange: skip testRange, as we already have all the ids
			//              |---------|
			//         |====================|
			if testRange[0] >= masterRange[0] && testRange[1] <= masterRange[1] {
				skip = true
				break
			}
			// case: testRange overlaps masterRange: replace with test range, as it has all the ids:
			// |----------------------------------|
			//         |====================|
			if testRange[0] < masterRange[0] && testRange[1] > masterRange[1] {
				nonOverlaps[j] = []int{-1, -1} // invalidate it
				// skip the rest of the tests, but keep running the other master tests:
				continue
			}
			// case: testRange overlaps masterRange on the left side: add the left overlap part to the overrides:
			// |----------------|
			//         |====================|
			if testRange[0] < masterRange[0] && testRange[1] >= masterRange[0] {
				testRange = []int{testRange[0], masterRange[0] - 1}
			}
			// case: testRange overlaps masterRange on the right side: add the right overlap part to the overrides:
			//             |----------------|
			// |====================|
			if testRange[0] <= masterRange[1] && testRange[1] > masterRange[1] {
				testRange = []int{masterRange[1] + 1, testRange[1]}
			}
		}
		if !skip {
			// ok, if we got so far, we do not overlap at all:
			nonOverlaps = append(nonOverlaps, testRange)
		}
	}
	// count all ranges in the overlap ranges:
	for _, r := range nonOverlaps {
		if r[0] == -1 {
			continue
		}
		d.s2 += (r[1] - r[0] + 1)
	}
}

func (d *Day05) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day05) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
