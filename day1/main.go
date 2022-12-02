package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"sort"
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
	// part 1
	a(input)
	// part 2
	b(input)
}

func a(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	max := 0
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if sum > max {
				max = sum
			}
			sum = 0

			continue
		}

		sum += utils.ToInt(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("result=", max)
}

func b(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var sums []int
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			sums = append(sums, sum)
			sum = 0
			continue
		}

		sum += utils.ToInt(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	result := 0
	for i := 0; i < 3; i++ {
		result += sums[i]
	}
	fmt.Println("result=", result)
}
