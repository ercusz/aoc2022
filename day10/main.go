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

	signals := make([][]string, len(lines))
	for idx, line := range lines {
		signals[idx] = append(signals[idx], strings.Split(line, " ")...)
	}
	return signals
}

func calSignalStrength(cycle int, x int) int {
	signalStrength := 0
	// 20th, 60th, 100th, 140th, 180th, and 220th
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		signalStrength = cycle * x
	}
	return signalStrength
}

func a(input string) {
	signals := parse(input)
	cycle := 0
	x := 1
	res := 0
	for _, signal := range signals {
		cycle++
		res += calSignalStrength(cycle, x)
		switch signal[0] {
		case "noop":
			continue
		case "addx":
			cycle++
			res += calSignalStrength(cycle, x)
			x += utils.ToInt(signal[1])
		}
	}
	fmt.Println(res)
}

func b(input string) {
	crtScreen := [6][]string{}
	for i := range crtScreen {
		for j := 0; j < 40; j++ {
			crtScreen[i] = append(crtScreen[i], ".")
		}
	}
	signals := parse(input)
	cycle := 0
	x := 1
	for _, signal := range signals {
		cycle++
		drawPixel(cycle, x, crtScreen)

		switch signal[0] {
		case "noop":
			continue
		case "addx":
			cycle++
			drawPixel(cycle, x, crtScreen)

			x += utils.ToInt(signal[1])
		}
	}

	for _, crt := range crtScreen {
		fmt.Println(strings.Join(crt, ""))
	}
}

func drawPixel(cycle int, x int, crtScreen [6][]string) {
	col := (cycle - 1) % 40
	row := ((cycle - 1) / 40) % 6

	if x-1 <= col && col <= x+1 {
		crtScreen[row][col] = "#"
	}
}
