package day12

import (
	"fmt"
	"regexp"
	"strings"

	"alexi.ch/aoc/2025/lib"
)

type Present [][]rune

func (p Present) String() string {
	str := "\n"
	for y := 0; y < len(p); y++ {
		for x := 0; x < len(p[y]); x++ {
			str += string(p[y][x])
		}
		str += "\n"
	}
	return str
}

func (p Present) UsedArea() int {
	sum := 0
	for y := 0; y < len(p); y++ {
		for x := 0; x < len(p[y]); x++ {
			if p[y][x] == '#' {
				sum += 1
			}
		}
	}
	return sum
}

type Area struct {
	width, height int
	presents      []int
}

type Day12 struct {
	s1       uint64
	s2       uint64
	presents []Present
	areas    []Area
}

func New() Day12 {
	return Day12{s1: 0, s2: 0}
}

func (d *Day12) Title() string {
	return "Day 12 - CHRISTMAS TREE FARM"
}

func (d *Day12) Setup() {
	// var lines = lib.ReadLines("data/12-test-data.txt")
	var lines = lib.ReadLines("data/12-data.txt")

	// parsing this is a real p.i.t.a, so I just hard-code it:
	// 6 presents with the number, 3 lines.
	for i := 0; i < 6; i++ {
		present := make([][]rune, 3)
		for y := 0; y < 3; y++ {
			present[y] = make([]rune, 3)
			for x := 0; x < 3; x++ {
				present[y][x] = rune(lines[i*5+y+1][x])
			}
		}
		d.presents = append(d.presents, present)
	}

	// from line 31ff, we have the area definitions:
	matcher := regexp.MustCompile(`(\d+)x(\d+): (.*)`)
	for i := 30; i < len(lines); i++ {
		line := lines[i]
		parts := matcher.FindStringSubmatch(line)
		if parts != nil {
			presentStrs := strings.Split(parts[3], " ")
			area := Area{
				width:    lib.StrToInt(parts[1]),
				height:   lib.StrToInt(parts[2]),
				presents: lib.Map(&presentStrs, func(s string) int { return lib.StrToInt(s) }),
			}
			d.areas = append(d.areas, area)
		}
	}

	// fmt.Printf("%#v\n", d.areas)
}

func (d *Day12) SolveProblem1() {
	d.s1 = 0
	for _, area := range d.areas {
		expectedArea := area.width * area.height
		presentArea := 0
		maxArea := 0
		for i, nr := range area.presents {
			presentArea += d.presents[i].UsedArea() * nr
			maxArea += 9 * nr
		}
		// easy case:
		// if the nr of pixels is larger than the area itself, it may never fit:
		if presentArea > expectedArea {
			// too large, skip
			continue
		}
		// easy case:
		// if we can align all 9x9 tiles without overlap, it will fit in any case:
		if maxArea <= expectedArea {
			d.s1++
			continue
		}

		// now we have to try to fit
		fmt.Printf("Area: %dx%d (%d): PresentArea: %d\n", area.width, area.height, expectedArea, presentArea)
		// Hmmmm..... That was unexpected.... there is NO problem with the real data that
		// needs fitting....??? let's see if that works....
		//
		//.... ok, that really was unexpected.... it worked without fitting at all.... just by checking
		// the two base cases above....
		// which is a bit annoying, since this only works with the final data, not with the test data.
	}
}

func (d *Day12) SolveProblem2() {
	d.s2 = 0
}

func (d *Day12) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day12) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
