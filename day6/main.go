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

func countCharUntilFindMarker(mSize int) int {
	chars := 0
	for idx := range input {
		if idx+mSize <= len(input) {
			marker := strings.Split(input[idx:idx+mSize], "")
			markerSet := utils.RemoveDuplicate(marker)
			if len(markerSet) == mSize {
				chars += idx + mSize
				break
			}
		}
	}
	return chars
}

func a(input string) {
	fmt.Println(countCharUntilFindMarker(4))
}

func b(input string) {
	fmt.Println(countCharUntilFindMarker(14))
}
