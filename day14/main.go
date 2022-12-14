package main

import (
	_ "embed"
	"fmt"
	"image"
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

func fillRockIntoSpace(input string) (map[image.Point]struct{}, int) {
	lines := strings.Split(input, "\n")

	spaces := map[image.Point]struct{}{}
	maxY := 0
	for _, line := range lines {
		paths := strings.Split(line, " -> ")
		for i := 0; i < len(paths)-1; i++ {
			var p1, p2 image.Point
			fmt.Sscanf(paths[i], "%d,%d", &p1.X, &p1.Y)
			fmt.Sscanf(paths[i+1], "%d,%d", &p2.X, &p2.Y)

			dx := utils.Sgn(p2.X - p1.X)
			dy := utils.Sgn(p2.Y - p1.Y)
			for j := (image.Point{dx, dy}); p1 != p2.Add(j); p1 = p1.Add(j) {
				spaces[p1] = struct{}{} // filled rock
				if p1.Y > maxY {
					maxY = p1.Y
				}
			}
		}
	}

	return spaces, maxY
}

func pourSand(spaces map[image.Point]struct{}, maxY int, countAbyss bool) int {
	sandUnits := 0
	for {
		p := image.Point{500, 0}
		for {
			n := p
			directions := []image.Point{{0, 1}, {-1, 1}, {1, 1}}
			for _, d := range directions {
				if _, isReserved := spaces[p.Add(d)]; !isReserved && p.Add(d).Y < maxY+2 {
					n = p.Add(d)
					break
				}
			}

			// if sand filled the whole space between rocks
			if countAbyss && n.Y == maxY {
				return sandUnits
			}

			if n == p {
				spaces[p] = struct{}{} // filled space
				sandUnits++
				break
			}
			p = n
		}
		if p.Y == 0 {
			break
		}
	}

	return sandUnits
}

func a(input string) {
	spaces, maxY := fillRockIntoSpace(input)
	res := pourSand(spaces, maxY, true)
	fmt.Println(res)
}

func b(input string) {
	spaces, maxY := fillRockIntoSpace(input)
	res := pourSand(spaces, maxY, false)
	fmt.Println(res)
}
