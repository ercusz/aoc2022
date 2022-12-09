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

func parse(input string) [][]string {
	lines := strings.Split(input, "\r\n")

	// no need to convert string to int
	moves := make([][]string, len(lines))
	for idx, line := range lines {
		moves[idx] = append(moves[idx], strings.Split(line, " ")...)
	}
	return moves
}

type pos struct {
	x, y int
}

func a(input string) {
	moves := parse(input)
	visited := make(map[pos]bool)
	head := pos{0, 0}
	tail := pos{0, 0}

	for _, move := range moves {
		for i := 0; i < utils.ToInt(move[1]); i++ {
			switch move[0] {
			case "U":
				head.y--
			case "D":
				head.y++
			case "L":
				head.x--
			case "R":
				head.x++
			}
			tail.moveTail(head)
			visited[tail] = true
		}
	}
	fmt.Println(len(visited))
}

func b(input string) {
	moves := parse(input)
	visited := make(map[pos]bool)
	knots := make([]pos, 10)
	visited[knots[9]] = true

	for _, move := range moves {
		for i := 0; i < utils.ToInt(move[1]); i++ {
			switch move[0] {
			case "U":
				knots[0].y--
			case "D":
				knots[0].y++
			case "L":
				knots[0].x--
			case "R":
				knots[0].x++
			}

			for k := range knots[:len(knots)-1] {
				knots[k+1].moveTail(knots[k])
			}
			visited[knots[9]] = true
		}

	}
	fmt.Println(len(visited))
}

func (t *pos) moveTail(h pos) {
	dy := t.y - h.y
	dx := t.x - h.x

	if utils.AbsInt(dx) < 2 && utils.AbsInt(dy) < 2 {
		return
	}

	if dy == 0 {
		if dx > 0 {
			t.x--
		} else {
			t.x++
		}
	} else if dx == 0 {
		if dy > 0 {
			t.y--
		} else {
			t.y++
		}
	} else {
		if dx > 0 {
			t.x--
		} else {
			t.x++
		}

		if dy > 0 {
			t.y--
		} else {
			t.y++
		}
	}
}
