package day10

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"

	"alexi.ch/aoc/2025/lib"
)

type Button struct {
	// the button's toggle map is a bitmap:
	// (1,3) means bits 1 and 3 need to be toggled.
	// We reverse the Lights array to use LSB direction:
	// [.  .  .  .] Lights array with 3 leds
	//  0  1  2  3 <-- bit nr
	// (1,3) means:
	// toggle from:
	// [.  .  .  .]
	// to:
	// [.  #  .  #]
	// so the toggleBitmap contains 1010 --> 10
	toggleBitmap uint64
}

type LightsState struct {
	// see description on Button:
	// the lights are a Bitmap, but reversed:
	// [#.#.] means: LED's 0 and 2 are on
	// --> to make bit operations simpler, store the bitmap reversed:
	// [#..#.] is stored as "01001" --> 9
	lightsBitmap uint64
	length       int

	neighbourEdges []*Edge

	// djikstra props:
	visited  bool
	distance int
}

// Input string is e.g.:
// .#.#
// output is
// lightsBitmap = 5 (0101, bit 2, 0 on)
func NewLightsStateFromInputStr(input string) LightsState {
	state := LightsState{length: len(input), distance: math.MaxInt, visited: false}
	for i, b := range input {
		switch b {
		case '.':
			// lightsMap = lightsMap << 1
		case '#':
			state.lightsBitmap = state.lightsBitmap | uint64(lib.PowInt(2, i))
		default:
			panic("Unsupported lights")
		}
	}
	return state
}

type Edge struct {
	btn     *Button
	toState *LightsState
}

type NodeMap map[byte]*LightsState

type InputData struct {
	targetLightsState LightsState
	buttons           []Button
	joltage           []int
}

type Day10 struct {
	s1        int
	s2        int
	nodeMap   NodeMap
	inputData []InputData
}

func New() Day10 {
	return Day10{s1: 0, s2: 0, nodeMap: make(NodeMap)}
}

func (d *Day10) Title() string {
	return "Day 10 - FACTORY"
}

func (d *Day10) Setup() {
	// var lines = lib.ReadLines("data/10-test-data.txt")
	var lines = lib.ReadLines("data/10-data.txt")
	lightsMatcher := regexp.MustCompile(`\[(.*)\]`)
	buttonsMatcher := regexp.MustCompile(`(\(([0-9,]+)\))+`)
	joltageMatcher := regexp.MustCompile(`\{(.*)\}`)

	for _, line := range lines {
		inputData := InputData{}

		// parse lights REVERSED into a bitmap represented by an int:
		// ints are LSB, our lights are MSB, so interpret the lights so that
		// LED 0 (left) is bit 0 (right)
		lights := lightsMatcher.FindStringSubmatch(line)
		if len(lights) == 2 {
			inputData.targetLightsState = NewLightsStateFromInputStr(lights[1])
			// fmt.Printf("%s --> %d\n", lights[1], inputData.targetLightsState.lightsBitmap)
		}

		// Button extraction
		buttonGroups := buttonsMatcher.FindAllStringSubmatch(line, -1)
		inputData.buttons = make([]Button, 0)
		for _, bgroup := range buttonGroups {
			button := bgroup[2]
			digitsStr := strings.Split(button, ",")
			digits := lib.Map(&digitsStr, func(s string) uint64 { return uint64(lib.StrToInt(s)) })
			var bitmap uint64 = 0
			for _, b := range digits {
				bitmap = bitmap | (1 << b)
			}
			// fmt.Printf("%s --> %d\n", button, bitmap)
			inputData.buttons = append(inputData.buttons, Button{toggleBitmap: bitmap})
		}

		// joltage extraction:
		jg := joltageMatcher.FindStringSubmatch(line)
		jstr := strings.Split(jg[1], ",")
		j := lib.Map(&jstr, func(i string) int { return lib.StrToInt(i) })
		inputData.joltage = j

		d.inputData = append(d.inputData, inputData)

	}
	// fmt.Printf("%v\n", d.numbers)

}

// This seems to be a Graph problem:
// the light states are vertices,
// while the buttons are the edges:
// so the button describes how to get from one state to another.
//
// we need to build all possible states (e.g. from [...] to [###]) (vertices),
// and apply all buttons on them to find the connected state.
// This forms a graph, which then can be traversed by a Shortest Path algo,
// e.g. with the beloved Djikstra!
func (d *Day10) SolveProblem1() {
	d.s1 = 0

	for _, input := range d.inputData {

		// fmt.Printf("%#v\n", input)

		// create light state map: create all possible states.
		// Those are the edges of the graph.
		lightsStates := make(map[uint64]*LightsState)
		for i := 0; i < lib.PowInt(2, input.targetLightsState.length); i++ {
			lightsStates[uint64(i)] = &LightsState{lightsBitmap: uint64(i), length: input.targetLightsState.length}
		}

		// for all lights states, create all possible neighbour states by applying all buttons.
		// So in the end, we have a map with all light states, with edges (buttons) to the next possible states
		for _, state := range lightsStates {
			for _, button := range input.buttons {
				nextLightBitmap := state.lightsBitmap ^ button.toggleBitmap
				nextLightState := lightsStates[nextLightBitmap]
				state.neighbourEdges = append(state.neighbourEdges, &Edge{btn: &button, toState: nextLightState})

			}
		}

		// now we have a complete lights state graph in lightsStates (vertices), which we can traverse
		// using Djikstra.
		// Built unvisited set:
		unvisited := make([]*LightsState, 0)
		var current *LightsState
		for _, state := range lightsStates {
			if state.lightsBitmap == 0 {
				// start node: make it the current one
				state.visited = true
				state.distance = 0
				current = state
			} else {
				// other nodes, mark as unvisited:
				state.visited = false
				state.distance = math.MaxInt
				unvisited = append(unvisited, state)
			}
		}

		// now, dance the djikstra!
		for {
			current.visited = true
			newDist := current.distance + 1
			// update unvisited neighbours:
			for _, n := range current.neighbourEdges {
				if !n.toState.visited && n.toState.distance > newDist {
					n.toState.distance = newDist
				}
			}
			// remove current from unvisited set:
			unvisited = lib.RemoveItem(unvisited, current)
			if len(unvisited) > 0 {
				// find smallest unvisited for next current:
				slices.SortFunc(unvisited, func(a, b *LightsState) int { return a.distance - b.distance })
				current = unvisited[0]
			} else {
				break
			}
		}
		// Distance of target node is the solution:
		dist := lightsStates[input.targetLightsState.lightsBitmap].distance
		// fmt.Printf("Dist: %d\n", dist)
		d.s1 += dist
	}

}

func (d *Day10) SolveProblem2() {
	d.s2 = 0
}

func (d *Day10) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day10) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
