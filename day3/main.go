package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

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

func calculatePriority(c string) int {
	r, _ := utf8.DecodeRuneInString(c) // convert string to rune
	res := 0
	if !unicode.IsLetter(r) {
		return 0
	}
	if unicode.IsLower(r) {
		res += int(1 + r - 'a')
	} else {
		res += int(27 + r - 'A')
	}

	return res
}

func a(input string) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		a := utils.RemoveDuplicate(strings.Split(line[:len(line)/2], ""))
		b := utils.RemoveDuplicate(strings.Split(line[len(line)/2:], ""))
		for _, c := range a {
			if strings.Contains(strings.Join(b, ""), c) {
				sum += calculatePriority(c)
			}
		}
	}
	fmt.Println(sum)
}

func b(input string) {
	input = strings.TrimRight(input, "\n")
	lines := strings.Split(input, "\n")
	sum := 0
	for i := 0; i < len(lines); i += 3 {
		l1 := utils.RemoveDuplicate(strings.Split(lines[i], ""))
		l2 := utils.RemoveDuplicate(strings.Split(lines[i+1], ""))
		l3 := utils.RemoveDuplicate(strings.Split(lines[i+2], ""))

		for _, c := range l1 {
			if strings.Contains(strings.Join(l2, ""), c) {
				if strings.Contains(strings.Join(l3, ""), c) {
					sum += calculatePriority(c)
				}
			}

		}
	}
	fmt.Println(sum)
}
