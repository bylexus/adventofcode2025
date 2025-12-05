package day04

import (
	"fmt"

	"alexi.ch/aoc/2025/lib"
)

type Day04 struct {
	s1   int
	s2   int
	grid map[lib.Coord]bool
}

func New() Day04 {
	return Day04{s1: 0, s2: 0, grid: make(map[lib.Coord]bool)}
}

func (d *Day04) Title() string {
	return "Day 04 - PRINTING DEPARTMENT"
}

func (d *Day04) Setup() {
	// var lines = lib.ReadLines("data/04-test-data.txt")
	var lines = lib.ReadLines("data/04-data.txt")
	for y, line := range lines {
		for x, char := range line {
			hasPaper := false
			if char == '@' {
				hasPaper = true
			}
			d.grid[lib.NewCoord2D(x, y)] = hasPaper
		}
	}
}

func (d *Day04) SolveProblem1() {
	d.s1 = 0
	dirs := lib.DIRS_2D_8_NEIGHBOURS

	for coord, val := range d.grid {
		if !val {
			continue
		}
		neighbours := 0
		for _, dir := range dirs {
			if d.grid[coord.Add(dir)] {
				neighbours++
			}
		}
		if neighbours < 4 {
			d.s1++
		}
	}
}

func (d *Day04) SolveProblem2() {
	d.s2 = 0
	dirs := lib.DIRS_2D_8_NEIGHBOURS
	grid := d.grid

	toBeRemoved := make([]lib.Coord, 0)

	for {
		for coord, val := range grid {
			if !val {
				continue
			}
			neighbours := 0
			for _, dir := range dirs {
				if grid[coord.Add(dir)] {
					neighbours++
				}
			}
			if neighbours < 4 {
				toBeRemoved = append(toBeRemoved, coord)
			}
		}
		if len(toBeRemoved) == 0 {
			break
		}
		d.s2 += len(toBeRemoved)
		for _, coord := range toBeRemoved {
			delete(grid, coord)
		}
		toBeRemoved = make([]lib.Coord, 0)
	}
}

func (d *Day04) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day04) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
