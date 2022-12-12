package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

type Point struct {
	x, y int
}

func main() {
	var start, end Point
	height := map[Point]rune{}

	// finding start and end points
	for x, s := range strings.Fields(input) {
		for y, r := range s {
			if r == 'S' {
				start = Point{x, y}
			} else if r == 'E' {
				end = Point{x, y}
			}
			height[Point{x, y}] = r
		}
	}

	height[start], height[end] = 'a', 'z'

	dist := map[Point]int{end: 0}
	queue := []Point{end}
	var shortest *Point
	directions := []Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // up, down, left, right

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if height[curr] == 'a' && shortest == nil {
			shortest = &curr
		}

		for _, near := range directions {
			next := Point{
				curr.x + near.x,
				curr.y + near.y,
			}
			_, seen := dist[next]
			_, valid := height[next]
			if seen || !valid || height[next] < height[curr]-1 {
				continue
			}

			dist[next] = dist[curr] + 1
			queue = append(queue, next)
		}
	}

	fmt.Println("part1:", dist[start])
	fmt.Println("part2:", dist[*shortest])
}
