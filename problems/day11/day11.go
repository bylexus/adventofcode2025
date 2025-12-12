package day11

import (
	"fmt"
	"regexp"
	"strings"

	"alexi.ch/aoc/2025/lib"
)

type Node struct {
	downstreamPaths int
	neighbours      []string
	name            string
}

type Day11 struct {
	s1    int
	s2    int
	graph map[string]*Node
}

func New() Day11 {
	return Day11{s1: 0, s2: 0, graph: make(map[string]*Node)}
}

func (d *Day11) Title() string {
	return "Day 11 - REACTOR"
}

func (d *Day11) Setup() {
	// var lines = lib.ReadLines("data/11-test-data.txt")
	// var lines = lib.ReadLines("data/11-test-data2.txt")
	var lines = lib.ReadLines("data/11-data.txt")
	matcher := regexp.MustCompile(`(.*):(.*)`)
	for _, line := range lines {
		matches := matcher.FindStringSubmatch(line)
		if len(matches) == 3 {
			nodeName := strings.TrimSpace(matches[1])
			neighbours := strings.Split(matches[2], " ")
			neighbours = lib.Map(&neighbours, func(s string) string { return strings.TrimSpace(s) })
			node := Node{
				name:            nodeName,
				downstreamPaths: -1,
				neighbours:      make([]string, 0),
			}
			for _, neighbour := range neighbours {
				neighbour = strings.TrimSpace(neighbour)
				if len(neighbour) > 0 {
					node.neighbours = append(node.neighbours, neighbour)
				}
			}
			d.graph[nodeName] = &node
		}
	}
	// fmt.Printf("%#v\n", d.graph["you"])
}

func (d *Day11) SolveProblem1() {
	d.s1 = 0
	d.s1 = d.countPathsToTarget("you", "out")
}

func (d *Day11) countPathsToTarget(start, end string) int {
	if start == end {
		return 1
	}
	if start == "out" {
		return 0
	}
	startNode := d.graph[start]
	if startNode.downstreamPaths >= 0 {
		return startNode.downstreamPaths
	}

	paths := 0
	for _, next := range startNode.neighbours {
		paths += d.countPathsToTarget(next, end)
	}
	startNode.downstreamPaths = paths
	return paths

}

func (d *Day11) SolveProblem2() {
	d.s2 = 0
	// Find all from svr to "fft"
	// Find all from fft to "dac"
	// Find all from dac to "out"
	// --> multiply

	// Find all from svr to "dac"
	// Find all from dac to "fft"
	// Find all from fft to "out"
	// --> multiply
	// sum both

	d.resetGraph()
	ret1 := d.countPathsToTarget("svr", "fft")
	d.resetGraph()
	ret2 := d.countPathsToTarget("fft", "dac")
	d.resetGraph()
	ret3 := d.countPathsToTarget("dac", "out")
	d.s2 = ret1 * ret2 * ret3

	d.resetGraph()
	ret1 = d.countPathsToTarget("svr", "dac")
	d.resetGraph()
	ret2 = d.countPathsToTarget("dac", "fft")
	d.resetGraph()
	ret3 = d.countPathsToTarget("fft", "out")
	d.s2 += ret1 * ret2 * ret3

}

func (d *Day11) resetGraph() {
	for _, n := range d.graph {
		n.downstreamPaths = -1
	}
}

func (d *Day11) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day11) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
