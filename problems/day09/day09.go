package day09

import (
	"fmt"
	"math"
	"strings"

	"alexi.ch/aoc/2025/lib"
)

type Day09 struct {
	s1    int
	s2    int
	tiles []lib.Coord
}

func New() Day09 {
	return Day09{s1: 0, s2: 0}
}

func (d *Day09) Title() string {
	return "Day 09 - MOVIE THEATER"
}

func (d *Day09) Setup() {
	var lines = lib.ReadLines("data/09-test-data.txt")
	// var lines = lib.ReadLines("data/09-data.txt")
	for _, line := range lines {
		splits := strings.Split(line, ",")
		if len(splits) == 2 {
			d.tiles = append(d.tiles, lib.NewCoord2D(
				lib.StrToInt(splits[0]),
				lib.StrToInt(splits[1]),
			))
		}
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day09) SolveProblem1() {
	d.s1 = 0
	maxArea := 0
	for i := 0; i < len(d.tiles)-1; i++ {
		for j := i + 1; j < len(d.tiles); j++ {
			a := d.tiles[i]
			b := d.tiles[j]
			area := (lib.Abs(a.X-b.X) + 1) * (lib.Abs(a.Y-b.Y) + 1)
			// fmt.Printf("A: %#v, B: %#v, area: %d\n", a, b, area)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	d.s1 = maxArea
}

func (d *Day09) SolveProblem2() {
	d.s2 = 0
	minCoord := lib.NewCoord2D(math.MaxInt, math.MaxInt)
	maxCoord := lib.NewCoord2D(math.MinInt, math.MinInt)
	tileMap := make(map[lib.Coord]byte)

	// calc outline:
	for i := 0; i < len(d.tiles); i++ {
		normI := i % len(d.tiles)
		normJ := (i + 1) % len(d.tiles)
		a := d.tiles[normI]
		b := d.tiles[normJ]
		// fmt.Printf("a: %#v, b: %#v\n", a, b)
		// draw line between 2 X:
		for x := lib.Min(a.X, b.X); x <= lib.Max(a.X, b.X); x++ {
			for y := lib.Min(a.Y, b.Y); y <= lib.Max(a.Y, b.Y); y++ {
				if minCoord.X > x {
					minCoord.X = x
				}
				if minCoord.Y > y {
					minCoord.Y = y
				}
				if maxCoord.X < x {
					maxCoord.X = x
				}
				if maxCoord.Y < y {
					maxCoord.Y = y
				}
				tileMap[lib.NewCoord2D(x, y)] = 'X'
			}
		}
		tileMap[a] = '#'
		tileMap[b] = '#'
	}

	// for y := minCoord.Y; y <= maxCoord.Y; y++ {
	// 	for x := minCoord.X; x <= maxCoord.X; x++ {
	// 		if entry, ok := tileMap[lib.NewCoord2D(x, y)]; ok {
	// 			fmt.Printf("%s", string(entry))
	// 		} else {
	// 			fmt.Printf("%s", ".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	// flood-fill - best guess is that the center of the area is within the border...
	queue := make([]lib.Coord, 0)
	center := lib.NewCoord2D((maxCoord.X+minCoord.X)/2, (maxCoord.Y+minCoord.Y)/2)
	fmt.Printf("min: %#v\n", minCoord)
	fmt.Printf("max: %#v\n", maxCoord)
	fmt.Printf("center: %#v\n", center)
	if _, ok := tileMap[center]; !ok {
		queue = append(queue, center)
	} else {
		panic("center tile is a border")
	}
	for len(queue) > 0 {
		coord := queue[0]
		queue = queue[1:]
		for _, d := range lib.DIRS_2D_8_NEIGHBOURS {
			n := coord.Add(d)
			// stop if we are out of bounds:
			if n.X < minCoord.X || n.X > maxCoord.X || n.Y < minCoord.Y || n.Y > maxCoord.Y {
				panic("We reached the border - we got the wrong start coordinate")
			}
			// continue flooding on an empty tile, stop on an already filled tile:
			if _, ok := tileMap[n]; !ok {
				tileMap[n] = 'X'
				queue = append(queue, n)
			}
		}
	}

	// ok, now check all the squares - that might take long...
	maxArea := 0
	for i := 0; i < len(d.tiles)-1; i++ {
	skip:
		for j := i + 1; j < len(d.tiles); j++ {
			a := d.tiles[i]
			b := d.tiles[j]
			// check the whole area if it is contained in tiles:
			for x := lib.Min(a.X, b.X); x <= lib.Max(a.X, b.X); x++ {
				for y := lib.Min(a.Y, b.Y); y <= lib.Max(a.Y, b.Y); y++ {
					if _, ok := tileMap[lib.NewCoord2D(x, y)]; !ok {
						continue skip
					}
				}
			}

			area := (lib.Abs(a.X-b.X) + 1) * (lib.Abs(a.Y-b.Y) + 1)
			// fmt.Printf("A: %#v, B: %#v, area: %d\n", a, b, area)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	d.s2 = maxArea

	// for y := minCoord.Y; y <= maxCoord.Y; y++ {
	// 	for x := minCoord.X; x <= maxCoord.X; x++ {
	// 		if entry, ok := tileMap[lib.NewCoord2D(x, y)]; ok {
	// 			fmt.Printf("%s", string(entry))
	// 		} else {
	// 			fmt.Printf("%s", ".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}

func (d *Day09) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day09) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
