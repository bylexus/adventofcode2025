package day02

import (
	"fmt"
	"strconv"
	"strings"

	"alexi.ch/aoc/2025/lib"
)

type Range struct {
	From int
	To   int
}

type Day02 struct {
	s1    int
	s2    int
	input []Range
}

func New() Day02 {
	return Day02{input: make([]Range, 0)}
}

func (d *Day02) Title() string {
	return "Day 02 - GIFT SHOP"
}

func (d *Day02) Setup() {
	// var lines = lib.ReadLines("data/02-test-data.txt")
	var lines = lib.ReadLines("data/02-data.txt")
	var pairs = strings.Split(lines[0], ",")

	for _, pair := range pairs {
		splitPair := strings.Split(pair, "-")
		d.input = append(d.input, Range{
			From: lib.StrToInt(splitPair[0]),
			To:   lib.StrToInt(splitPair[1]),
		})
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day02) SolveProblem1() {
	d.s1 = 0
	for _, pair := range d.input {
		for i := pair.From; i <= pair.To; i++ {
			str := strconv.Itoa(i)
			lngth := len(str)
			mid := lngth / 2
			if lngth%2 == 0 && str[:mid] == str[mid:] {
				d.s1 += i
			}
		}
	}
}

func (d *Day02) SolveProblem2() {
	d.s2 = 0
	for _, pair := range d.input {
		for i := pair.From; i <= pair.To; i++ {
			str := strconv.Itoa(i)
			for ln := 1; ln <= len(str)/2; ln++ {
				splits := lib.CutIntoPartsOfLen(str, ln)
				splitMap := make(map[string]int)
				for _, split := range splits {
					splitMap[split] += 1
				}
				if len(splitMap) == 1 && splitMap[splits[0]] > 1 {
					d.s2 += i
					// fmt.Printf("Invalid: %d\n", i)
					// fmt.Printf("  splits: %#v\n", splits)
					break
				}
			}
		}
	}
}

func (d *Day02) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day02) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
