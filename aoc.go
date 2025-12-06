package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"alexi.ch/aoc/2025/problems"
	"alexi.ch/aoc/2025/problems/day01"

	//template:"alexi.ch/aoc/2025/problems/day{{- .ProblemNumber | format "%02d" -}}"
"alexi.ch/aoc/2025/problems/day06"
"alexi.ch/aoc/2025/problems/day05"
"alexi.ch/aoc/2025/problems/day04"
	"alexi.ch/aoc/2025/problems/day02"
	"alexi.ch/aoc/2025/problems/day03"
)

func main() {
	tannenbaum()
	var problem_map = map[string](func() problems.Problem){
		"01": func() problems.Problem { p := day01.New(); return &p },
		//template:"{{- .ProblemNumber | format "%02d" -}}": func() problems.Problem { p := day{{- .ProblemNumber | format "%02d" -}}.New(); return &p },
"06": func() problems.Problem { p := day06.New(); return &p },
"05": func() problems.Problem { p := day05.New(); return &p },
"04": func() problems.Problem { p := day04.New(); return &p },
		"03":         func() problems.Problem { p := day03.New(); return &p },
		"02":         func() problems.Problem { p := day02.New(); return &p },
		"playground": func() problems.Problem { p := problems.NewPlayground(); return &p },
	}

	var to_solve = os.Args[1:]

	if len(to_solve) == 0 {
		var keys = make([]string, 0)
		for key := range problem_map {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		to_solve = keys
	}

	// Run solving all problems  in parallel:
	var start = time.Now()
	var wg sync.WaitGroup
	wg.Add(len(to_solve))
	for _, p := range to_solve {
		go func(probKey string) {
			defer wg.Done()
			var prob = problem_map[probKey]
			if prob != nil {
				problems.Solve(prob())
			} else {
				panic("Problem not found")
			}
		}(p)
	}
	wg.Wait()
	var duration = time.Since(start)
	fmt.Printf("\n\nFull runtime: %s\n\n", duration)
}

func tannenbaum() {
	var t = strings.Join([]string{
		"\x1B[1;97m",
		"Advent of Code 2025",
		"--------------------",
		"",
		"        \x1B[1;93m*   *",
		"         \\ /",
		"         AoC",
		"         -\x1B[1;91m*\x1B[1;93m-",
		"          \x1B[1;37m|\x1B[0;32m",
		"          *",
		"         /*\\",
		"        /\x1B[1;94m*\x1B[0;32m*\x1B[1;93m*\x1B[0;32m\\",
		"       /\x1B[1;91m*\x1B[0;32m***\x1B[1;94m*\x1B[0;32m\\",
		"      /**\x1B[1;93m*\x1B[0;32m****\\",
		"     /**\x1B[1;94m*\x1B[0;32m***\x1B[1;91m*\x1B[0;32m**\\",
		"    /********\x1B[1;93m*\x1B[0;32m**\\",
		"   /**\x1B[1;91m*\x1B[0;32m*****\x1B[1;94m*\x1B[0;32m****\\",
		"  /**\x1B[1;94m*\x1B[0;32m*\x1B[1;93m*\x1B[0;32m**********\\",
		" /**\x1B[1;94m*\x1B[0;32m*****\x1B[1;93m*\x1B[0;32m**\x1B[1;91m*\x1B[0;32m****\x1B[1;93m*\x1B[0;32m\\",
		"          #",
		"          #",
		"       \x1B[1;97m2-0-2-5",
		"       #######",
		"\x1B[0m",
	}, "\n")
	fmt.Print(t)
}
