package lib

import "fmt"

type Coord struct {
	X int
	Y int
	Z int
}

func NewCoord3D(x int, y int, z int) Coord {
	return Coord{X: x, Y: y, Z: z}
}

func NewCoord2D(x int, y int) Coord {
	return NewCoord3D(x, y, 0)
}

func NewCoord0() Coord {
	return NewCoord3D(0, 0, 0)
}

func (c Coord) String() string {
	return fmt.Sprintf("{x: %d, y: %d, z: %d}", c.X, c.Y, c.Z)
}

func (c Coord) Manhattan(o Coord) int {
	return Abs(c.X-o.X) + Abs(c.Y-o.Y) + Abs(c.Z-o.Z)
}

func (c Coord) Add(o Coord) Coord {
	return NewCoord3D(c.X+o.X, c.Y+o.Y, c.Z+o.Z)
}

func (c Coord) AddXY(x, y int) Coord {
	return NewCoord3D(c.X+x, c.Y+y, c.Z)
}
