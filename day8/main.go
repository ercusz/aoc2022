package main

import (
	_ "embed"
	"fmt"
	"sort"
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

func main() {
	a(input)
	b(input)
}

func parse(input string) [][]string {
	lines := strings.Split(input, "\r\n")

	// no need to convert string to int
	grid := make([][]string, len(lines))
	for idx, line := range lines {
		grid[idx] = append(grid[idx], strings.Split(line, "")...)
	}
	return grid
}

func counting(grid [][]string) (int, int) {
	height, width := len(grid), len(grid[0])
	count := 2*width + 2*height - 4 // remove edge and then use to count for part 1
	scenics := []int{}              // use to store scenics score for part 2

	for r := 1; r < height-1; r++ {
		row := grid[r]
		for c := 1; c < height-1; c++ {
			col := row[c]

			hor, ver := 2, 2
			left, right, top, bottom := 0, 0, 0, 0

			// from col -> right
			for i := c + 1; i < len(row); i++ {
				right++
				if row[i] >= col {
					hor--
					break
				}
			}

			// from col -> left
			for i := c - 1; i >= 0; i-- {
				left++
				if row[i] >= col {
					hor--
					break
				}
			}

			// from col -> top
			for i := r + 1; i < len(grid); i++ {
				top++
				if grid[i][c] >= col {
					ver--
					break
				}
			}

			// from col -> bottom
			for i := r - 1; i >= 0; i-- {
				bottom++
				if grid[i][c] >= col {
					ver--
					break
				}
			}

			if hor > 0 || ver > 0 {
				count++
				scenics = append(scenics, left*right*top*bottom)
			}
		}
	}

	// get highest scenic view
	sort.Sort(sort.Reverse(sort.IntSlice(scenics))) // sort in descending order
	max := scenics[0]

	return count, max
}

func a(input string) {
	grid := parse(input)
	ans, _ := counting(grid)
	fmt.Println(ans)
}

func b(input string) {
	grid := parse(input)
	_, ans := counting(grid)
	fmt.Println(ans)
}
