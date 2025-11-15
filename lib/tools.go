package lib

import (
	"bufio"
	"errors"
	"os"
	"slices"
	"strconv"
	"sync"

	"github.com/bylexus/go-stdlib/eerr"
	"golang.org/x/exp/constraints"
)

var MOVE_VEC_2D_4DIRS = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLines(path string) []string {
	f, err := os.Open(path)
	Check(err)
	defer f.Close()

	var lines = make([]string, 0)

	r := bufio.NewScanner(f)
	r.Split(bufio.ScanLines)
	for r.Scan() {
		lines = append(lines, r.Text())
	}
	return lines
}

func FindMax[V constraints.Ordered](slice []V) (*V, error) {
	if len(slice) == 0 {
		return nil, errors.New("Empty slice")
	}
	var max V = slice[0]
	for i, v := range slice {
		if i == 0 || v > max {
			max = v
		}
	}
	return &max, nil
}

func Sum[V constraints.Integer](slice []V) V {
	var s V = 0
	for _, n := range slice {
		s += n
	}
	return s
}

/**
 * map function for slice: maps slice[I] to slice[R] by
 * using f(T) R
 */
func Map[I any, R any](s *[]I, f func(item I) R) []R {
	var result = make([]R, 0, len(*s))
	for _, item := range *s {
		result = append(result, f(item))
	}
	return result
}

func Max[T constraints.Ordered](a T, b T) T {
	if a >= b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a T, b T) T {
	if a <= b {
		return a
	}
	return b
}

func Abs[T constraints.Signed](a T) T {
	if a < 0 {
		return -1 * a
	}
	return a
}

func Splice[T any](slice []T, index int) []T {
	newSlice := make([]T, 0, len(slice)-1)
	for i := 0; i < len(slice); i++ {
		if i != index {
			newSlice = append(newSlice, slice[i])
		}
	}
	return newSlice
}

func Contains[T comparable](list []T, el T) bool {
	for _, a := range list {
		if a == el {
			return true
		}
	}
	return false
}

// greatest common divisor (GCD) via Euclidean algorithm
// source: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
// source: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func StrToInt(s string) int {
	n, err := strconv.Atoi(s)
	eerr.PanicOnErr(err)
	return n
}

func StrToUint64(s string) uint64 {
	n, err := strconv.ParseUint(s, 10, 64)
	eerr.PanicOnErr(err)
	return n
}

// Returns all permutations of a given set of values, for a given length.
// For example all combinations of "A" and "B" for a given length of 3 (e.g. "AAA", "ABA" etc.)
//
// Example:
//
//	perms := Permutations([]string{"A", "B"}, 3)
//	// -> [["A", "A", "A"], ["A", "A", "B"], ["A", "B", "A"], ["A", "B", "B"], ["B", "A", "A"], ["B", "A", "B"], ["B", "B", "A"], ["B", "B", "B"]]
func Permutations[T any](values []T, length int) [][]T {
	res := make([][]T, 0)
	if length <= 0 {
		return res
	}

	singlePerms := make([]T, 0, len(values))
	singlePerms = slices.Concat(singlePerms, values)

	if length == 1 {
		for _, perm := range values {
			res = append(res, []T{perm})
		}
	} else {
		prevPerms := Permutations(values, length-1)
		for _, perm := range singlePerms {
			for _, prevPerm := range prevPerms {
				newPerm := slices.Concat(prevPerm, []T{perm})
				res = append(res, newPerm)
			}
		}
	}

	return res
}

// Memoization wrapper for a function with 1 argument and 1 return value:
// It returns a function with the same signature as the given function.
// If this function is called, the result will be memoized, so subsequent calls to
// the function with ther same argument returns a cached result.
//
// The function is also thread-safe, so you can use it in a goroutine context.
func Memoize[T comparable, U any](fn func(a T) U) func(a T) U {
	cache := make(map[T]U)
	mutex := sync.Mutex{}
	return func(a T) U {
		key := a
		mutex.Lock()
		if v, ok := cache[key]; ok {
			mutex.Unlock()
			return v
		}
		mutex.Unlock()
		res := fn(a)
		mutex.Lock()
		cache[key] = res
		mutex.Unlock()
		return res
	}
}
