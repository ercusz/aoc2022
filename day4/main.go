package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/ercusz/aoc2022/utils"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	a(input)
	b(input)
}

func parse(input string) (res [][]int) {
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		a := strings.Split(pairs[0], "-")
		b := strings.Split(pairs[1], "-")
		res = append(res, []int{
			utils.ToInt(a[0]), utils.ToInt(a[1]),
			utils.ToInt(b[0]), utils.ToInt(b[1]),
		})
	}

	return
}

func isInRange(a, b []int) bool {
	return a[0] >= b[0] && a[1] <= b[1]
}

func isOverlap(a, b []int) bool {
	// swap if a > b (order lower to bigger)
	if a[0] > b[1] {
		a, b = b, a
	}
	// a : 1 2 3
	// b :   2 3 4 5
	return a[1] >= b[0]
}

func a(input string) {
	lines := parse(input)
	count := 0
	for _, line := range lines {
		a := line[:2]
		b := line[2:]
		if isInRange(a, b) || isInRange(b, a) {
			count++
		}
	}
	fmt.Println(count)
}

func b(input string) {
	lines := parse(input)
	count := 0
	for _, line := range lines {
		a := line[:2]
		b := line[2:]
		if isOverlap(a, b) {
			count++
		}
	}
	fmt.Println(count)
}
