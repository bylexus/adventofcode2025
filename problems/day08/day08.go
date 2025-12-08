package day08

import (
	"fmt"
	"math"
	"regexp"
	"slices"

	"alexi.ch/aoc/2025/lib"
)

type JunctionBox struct {
	Coord   lib.Coord
	Circuit int
}

type Pair struct {
	Boxes    []*JunctionBox
	Distance float64
}

func (p *Pair) First() *JunctionBox {
	return p.Boxes[0]
}
func (p *Pair) Second() *JunctionBox {
	return p.Boxes[1]
}

func (p *Pair) String() string {
	out := fmt.Sprintf("%s :: %s --> %f", p.Boxes[0], p.Boxes[1], p.Distance)

	return out
}

func (j *JunctionBox) String() string {
	out := fmt.Sprintf("%s (C: %d)", j.Coord.String(), j.Circuit)
	// if j.Nearest != nil {
	// 	out += fmt.Sprintf(", nearest: %d: %s", j.DistToNearest, j.Nearest.String())
	// }

	return out
}

type Day08 struct {
	s1         int
	s2         int
	boxes      []*JunctionBox
	circuitMap map[int][]*JunctionBox
}

func New() Day08 {
	return Day08{s1: 0, s2: 0, boxes: make([]*JunctionBox, 0), circuitMap: make(map[int][]*JunctionBox)}
}

func (d *Day08) Title() string {
	return "Day 08 - PLAYGROUND"
}

func (d *Day08) Setup() {
	// var lines = lib.ReadLines("data/08-test-data.txt")
	var lines = lib.ReadLines("data/08-data.txt")
	matcher := regexp.MustCompile(`(\d+),(\d+),(\d+)`)
	for _, line := range lines {
		parts := matcher.FindStringSubmatch(line)
		if parts != nil {
			b := JunctionBox{
				Coord: lib.NewCoord3D(
					lib.StrToInt(parts[1]),
					lib.StrToInt(parts[2]),
					lib.StrToInt(parts[3]),
				),
			}
			d.boxes = append(d.boxes, &b)
		}
	}
}

func (d *Day08) SolveProblem1() {
	d.s1 = 0
	// form all possible pairs of JunctionBoxes
	pairs := make([]Pair, 0)
	for i := 0; i < len(d.boxes)-1; i++ {
		for j := i + 1; j < len(d.boxes); j++ {
			a := d.boxes[i]
			b := d.boxes[j]
			a.Circuit = 0
			b.Circuit = 0
			dist := distance(a, b)
			pair := Pair{
				Boxes:    []*JunctionBox{a, b},
				Distance: dist,
			}
			pairs = append(pairs, pair)
		}
	}

	// sort pairs by shortest distances:
	slices.SortFunc(pairs, func(a, b Pair) int {
		dist := a.Distance - b.Distance
		if dist < 0 {
			return -1
		} else {
			return 1
		}
	})

	// connect to circuits:
	circuitCounter := 1
	connectionCounter := 0
	maxConnections := 1000
	for _, p := range pairs {
		if connectionCounter >= maxConnections {
			break
		}
		// both boxes unconnected? connect to a new circuit:
		if p.First().Circuit == 0 && p.Second().Circuit == 0 {
			p.First().Circuit = circuitCounter
			p.Second().Circuit = circuitCounter
			d.circuitMap[circuitCounter] = append(d.circuitMap[circuitCounter], p.First())
			d.circuitMap[circuitCounter] = append(d.circuitMap[circuitCounter], p.Second())
			circuitCounter++
		} else if p.First().Circuit > 0 && p.Second().Circuit == 0 {
			// 1st is connected, 2nd not: add 2nd to 1st:
			p.Second().Circuit = p.First().Circuit
			d.circuitMap[p.First().Circuit] = append(d.circuitMap[p.First().Circuit], p.Second())
		} else if p.First().Circuit == 0 && p.Second().Circuit > 0 {
			// 2nd is connected, 1st not: add 1st to 2nd:
			p.First().Circuit = p.Second().Circuit
			d.circuitMap[p.Second().Circuit] = append(d.circuitMap[p.Second().Circuit], p.First())
		} else if p.First().Circuit == p.Second().Circuit {
			// both boxes connected to the same circuit? nothing to be done
		} else if p.First().Circuit != p.Second().Circuit {
			// both boxes connected, but not to the same circuit:
			// move all from 2nd to 1st:
			toDel := p.Second().Circuit
			for _, b := range d.circuitMap[toDel] {
				b.Circuit = p.First().Circuit
				d.circuitMap[p.First().Circuit] = append(d.circuitMap[p.First().Circuit], b)
			}
			delete(d.circuitMap, toDel)
		} else {
			panic("cannot happen")
		}
		connectionCounter++
	}

	// count largest 3 circuits:
	circuitLengths := []int{}
	for _, boxes := range d.circuitMap {
		circuitLengths = append(circuitLengths, len(boxes))
	}
	slices.Sort(circuitLengths)
	slices.Reverse(circuitLengths)
	longest3 := circuitLengths[:3]
	d.s1 = 1
	for _, l := range longest3 {
		d.s1 *= l
	}
}

func (d *Day08) SolveProblem2() {
	d.s2 = 0
	d.circuitMap = make(map[int][]*JunctionBox)

	// form all possible pairs of JunctionBoxes
	pairs := make([]Pair, 0)
	for i := 0; i < len(d.boxes)-1; i++ {
		for j := i + 1; j < len(d.boxes); j++ {
			a := d.boxes[i]
			b := d.boxes[j]
			a.Circuit = 0
			b.Circuit = 0
			dist := distance(a, b)
			pair := Pair{
				Boxes:    []*JunctionBox{a, b},
				Distance: dist,
			}
			pairs = append(pairs, pair)
		}
	}

	// sort pairs by shortest distances:
	slices.SortFunc(pairs, func(a, b Pair) int {
		dist := a.Distance - b.Distance
		if dist < 0 {
			return -1
		} else {
			return 1
		}
	})

	// connect to circuits:
	circuitCounter := 1
	rounds := 0
	minRounds := 1000
	var lastPair Pair
outer:
	for {
		for _, p := range pairs {
			// both boxes unconnected? connect to a new circuit:
			if p.First().Circuit == 0 && p.Second().Circuit == 0 {
				p.First().Circuit = circuitCounter
				p.Second().Circuit = circuitCounter
				d.circuitMap[circuitCounter] = append(d.circuitMap[circuitCounter], p.First())
				d.circuitMap[circuitCounter] = append(d.circuitMap[circuitCounter], p.Second())
				circuitCounter++
			} else if p.First().Circuit > 0 && p.Second().Circuit == 0 {
				// 1st is connected, 2nd not: add 2nd to 1st:
				p.Second().Circuit = p.First().Circuit
				d.circuitMap[p.First().Circuit] = append(d.circuitMap[p.First().Circuit], p.Second())
			} else if p.First().Circuit == 0 && p.Second().Circuit > 0 {
				// 2nd is connected, 1st not: add 1st to 2nd:
				p.First().Circuit = p.Second().Circuit
				d.circuitMap[p.Second().Circuit] = append(d.circuitMap[p.Second().Circuit], p.First())
			} else if p.First().Circuit == p.Second().Circuit {
				// both boxes connected to the same circuit? nothing to be done
			} else if p.First().Circuit != p.Second().Circuit {
				// both boxes connected, but not to the same circuit:
				// move all from 2nd to 1st:
				toDel := p.Second().Circuit
				for _, b := range d.circuitMap[toDel] {
					b.Circuit = p.First().Circuit
					d.circuitMap[p.First().Circuit] = append(d.circuitMap[p.First().Circuit], b)
				}
				delete(d.circuitMap, toDel)
			} else {
				panic("cannot happen")
			}

			// check if we only have 1 circuit left, with no lose Junction Boxes:
			if len(d.circuitMap) == 1 && rounds > minRounds {
				allDone := true
				for _, b := range d.boxes {
					if b.Circuit == 0 {
						allDone = false
						break
					}
				}
				if allDone {
					lastPair = p
					break outer
				}
			}
			rounds++
		}
	}

	// fmt.Printf("Last pair: %s\n", &lastPair)
	d.s2 = lastPair.First().Coord.X * lastPair.Second().Coord.X
}

func (d *Day08) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day08) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

func distance(a, b *JunctionBox) float64 {
	return math.Sqrt(
		float64((a.Coord.X-b.Coord.X)*(a.Coord.X-b.Coord.X)) +
			float64((a.Coord.Y-b.Coord.Y)*(a.Coord.Y-b.Coord.Y)) +
			float64((a.Coord.Z-b.Coord.Z)*(a.Coord.Z-b.Coord.Z)),
	)
}
