package day07

import (
	"fmt"

	"alexi.ch/aoc/2025/lib"
)

type FieldMap map[lib.Coord]rune

type Field struct {
	FieldMap
	Width  int
	Height int
}

func (f Field) String() string {
	res := ""
	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			res += string(f.FieldMap[lib.NewCoord2D(x, y)])
		}
		res += "\n"
	}
	return res
}

type Day07 struct {
	s1             int
	s2             int
	field          Field
	field2         Field
	heads          []lib.Coord
	start2         lib.Coord
	timelineMemory map[lib.Coord]int
}

func New() Day07 {
	return Day07{s1: 0, s2: 0, field: Field{FieldMap: make(FieldMap)}, field2: Field{FieldMap: make(FieldMap)}, heads: make([]lib.Coord, 0), timelineMemory: make(map[lib.Coord]int)}
}

func (d *Day07) Title() string {
	return "Day 07 - LABORATORIES"
}

func (d *Day07) Setup() {
	// var lines = lib.ReadLines("data/07-test-data.txt")
	var lines = lib.ReadLines("data/07-data.txt")
	for y, line := range lines {
		d.field.Height = lib.Max(d.field.Height, y+1)
		d.field2.Height = lib.Max(d.field2.Height, y+1)
		for x, chr := range line {
			d.field.Width = lib.Max(d.field.Width, x+1)
			d.field2.Width = lib.Max(d.field2.Width, x+1)
			coord := lib.NewCoord2D(x, y)
			switch chr {
			case 'S':
				d.heads = append(d.heads, coord)
				d.start2 = coord
				d.field.FieldMap[coord] = '|'
				d.field2.FieldMap[coord] = '|'
			case '.':
				d.field.FieldMap[coord] = '.'
				d.field2.FieldMap[coord] = '.'
			case '^':
				d.field.FieldMap[coord] = '^'
				d.field2.FieldMap[coord] = '^'
			default:
				panic("Should not happen")
			}
		}
	}
}

func (d *Day07) SolveProblem1() {
	d.s1 = 0
	// fmt.Println(d.field)
	// fmt.Printf("%v\n", d.heads)
	splitCounter := 0

	for len(d.heads) > 0 {
		newHeads := make([]lib.Coord, 0)
		for _, h := range d.heads {
			nextCoord := h.AddXY(0, 1)
			nextField := d.field.FieldMap[nextCoord]

			// decide where the head of the beam goes, or what happens with it:
			switch nextField {
			case '.':
				// below is an empty space, so move ahead
				d.field.FieldMap[nextCoord] = '|'
				newHeads = append(newHeads, nextCoord)
			case '|':
				// do nothing: if this ever happens, the beam is "eaten"
			case '^':
				// split head into two new heads, check if there is already a beam there:
				splitCounter++
				leftCoord := nextCoord.AddXY(-1, 0)
				rightCoord := nextCoord.AddXY(1, 0)
				if d.field.FieldMap[leftCoord] == '.' {
					d.field.FieldMap[leftCoord] = '|'
					newHeads = append(newHeads, leftCoord)
				}
				if d.field.FieldMap[rightCoord] == '.' {
					d.field.FieldMap[rightCoord] = '|'
					newHeads = append(newHeads, rightCoord)
				}
			}
		}
		d.heads = newHeads
		// fmt.Println(d.field.String() + "\n")
		// fmt.Printf("%v\n", d.heads)
	}
	d.s1 = splitCounter
}

func (d *Day07) SolveProblem2() {
	d.s2 = 0
	d.s2 = d.CountTimelinesFrom(d.field2, d.start2)
}

func (d *Day07) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day07) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}

// Recursively travel down, count the possible timelines "below" the actual beam,
// while memoizing the already counted timelines.
func (d *Day07) CountTimelinesFrom(field Field, start lib.Coord) int {
	nextCoord := start.AddXY(0, 1)
	if timelines, ok := d.timelineMemory[nextCoord]; ok {
		return timelines
	}
	if next, ok := field.FieldMap[nextCoord]; ok {
		switch next {
		case '.':
			// next field is empty: just pass it down:
			ret := d.CountTimelinesFrom(field, nextCoord)
			d.timelineMemory[start] = ret
			return ret
		case '^':
			// start 2 timelines, left and right of the splitter:
			retLeft := d.CountTimelinesFrom(field, nextCoord.AddXY(-1, 0))
			retRight := d.CountTimelinesFrom(field, nextCoord.AddXY(1, 0))
			d.timelineMemory[start] = retLeft + retRight
			return retLeft + retRight
		default:
			panic("oops, should not happen")
		}
	} else {
		// if the next coord does not exist, we reached the end: return 1 timeline
		d.timelineMemory[start] = 1
		return 1
	}
}
