package problems

import (
	"fmt"
)

type Playground struct {
}

func NewPlayground() Playground {
	return Playground{}
}

func (d *Playground) Title() string {
	return "Playground"
}

func (d *Playground) Setup() {
}

func (d *Playground) SolveProblem1() {
	// fn := lib.Memoize(func(a int) int { return 5 })
	fmt.Printf("-5 %% 2 == %d\n", -5%2)
}

func (d *Playground) SolveProblem2() {
}

func (d *Playground) Solution1() string {
	return fmt.Sprintf("%d", 0)
}

func (d *Playground) Solution2() string {
	return fmt.Sprintf("%d", 0)
}

func modSlice(a []int, idx int, value int) {
	a[idx] = value
}
