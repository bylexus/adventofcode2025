{{- $nrStr := .ProblemNumber | format "%02d" -}}
package day{{- $nrStr }}

import (
	"fmt"

	"alexi.ch/aoc/2025/lib"
)

type Day{{- $nrStr }} struct {
	s1 uint64
	s2 uint64
}

func New() Day{{- $nrStr }} {
	return Day{{- $nrStr }}{s1: 0, s2: 0}
}

func (d *Day{{- $nrStr }}) Title() string {
	return "Day {{ $nrStr }} - {{ .Title }}"
}

func (d *Day{{- $nrStr -}}) Setup() {
	// var lines = lib.ReadLines("data/{{- $nrStr -}}-test-data.txt")
	var lines = lib.ReadLines("data/{{- $nrStr -}}-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day{{- $nrStr -}}) SolveProblem1() {
	d.s1 = 0
}

func (d *Day{{- $nrStr -}}) SolveProblem2() {
	d.s2 = 0
}

func (d *Day{{- $nrStr -}}) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day{{- $nrStr -}}) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
