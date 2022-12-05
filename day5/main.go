package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"unicode"

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

func getStacks(input string) [][]string {
	lines := strings.Split(input, "\r\n")
	numOfStacks := int(math.Ceil(float64(len(lines[0])) / 4))
	stacks := make([][]string, numOfStacks)
	// ex round = 5 | input round = 8
	for i := 0; i < 8; i++ {
		for idx, c := range lines[i] {
			if string(c) != "[" {
				continue
			}

			desIdx := idx / 4
			if unicode.IsLetter(rune(lines[i][idx+1])) {
				stacks[desIdx] = append(stacks[desIdx], string(lines[i][idx+1]))
			}
		}
	}

	return stacks
}

func getMove(input string) (moves [][]int) {
	lines := strings.Split(input, "\r\n")
	// ex i=5 | input i= 10
	for i := 10; i < len(lines); i++ {
		line := strings.ReplaceAll(lines[i], "move ", "")
		line = strings.ReplaceAll(line, " from ", " ")
		line = strings.ReplaceAll(line, " to ", " ")
		move := strings.Split(line, " ")
		moveInt := []int{}
		for _, m := range move {
			moveInt = append(moveInt, utils.ToInt(m))
		}
		moves = append(moves, moveInt)
	}

	return
}

func prepend(newData []string, origin []string, isReversed bool) []string {
	res := []string{}
	if isReversed {
		utils.Reverse(newData)
	}
	res = append(res, newData...)
	res = append(res, origin...)
	return res
}

func printResult(stacks [][]string) {
	result := []string{}
	for _, s := range stacks {
		result = append(result, s[0])
	}
	fmt.Println("ans=", strings.Join(result, ""))
}

func a(input string) {
	stacks := getStacks(input)
	moves := getMove(input)

	for _, move := range moves {
		numOfCrates := move[0]
		from := move[1] - 1
		to := move[2] - 1

		// fmt.Printf("move %v from %v to %v\n", stacks[from][:numOfCrates], stacks[from], stacks[to])
		stacks[to] = prepend(stacks[from][:numOfCrates], stacks[to], true)
		stacks[from] = stacks[from][numOfCrates:]
	}

	printResult(stacks)
}

func b(input string) {
	stacks := getStacks(input)
	moves := getMove(input)

	for _, move := range moves {
		numOfCrates := move[0]
		from := move[1] - 1
		to := move[2] - 1

		// fmt.Printf("move %v from %v to %v\n", stacks[from][:numOfCrates], stacks[from], stacks[to])
		stacks[to] = prepend(stacks[from][:numOfCrates], stacks[to], false)
		stacks[from] = stacks[from][numOfCrates:]
	}

	printResult(stacks)
}
