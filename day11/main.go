package main

import (
	_ "embed"
	"encoding/json"
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

type Monkey struct {
	items     []int
	operation func(int) int
	test      func(int) int
}

const monkeyTemplate = `Monkey %d:
  Starting items: %s
  Operation: new = old %s %d
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d`

func parse(input string) ([]Monkey, int) {
	notes := strings.Split(input, "\n\n")
	monkeys, lcm := make([]Monkey, len(notes)), 1
	for _, s := range notes {
		var items, op string
		var i, v, test, t, f int

		// remove space in "items" line to use json.Unmarshal easier
		// replace "old * old" with "^ 2" to follow the same format as other operations
		s = strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(s)

		fmt.Sscanf(s, monkeyTemplate, &i, &items, &op, &v, &test, &t, &f)

		json.Unmarshal([]byte("["+items+"]"), &monkeys[i].items)

		monkeys[i].operation = map[string]func(int) int{
			"+": func(o int) int { return o + v },
			"*": func(o int) int { return o * v },
			"^": func(o int) int { return o * o },
		}[op]

		monkeys[i].test = func(w int) int {
			if w%test == 0 {
				return t
			}
			return f
		}
		lcm *= test
	}

	return monkeys, lcm
}

func inspect(monkeys []Monkey, rounds int, operation func(int) int) int {
	inspects := make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for j, monkey := range monkeys {
			for _, item := range monkey.items {
				worry := operation(monkey.operation(item))

				receiverMonkey := &monkeys[monkey.test(worry)]
				receiverMonkey.items = append(receiverMonkey.items, worry)

				inspects[j]++
			}
			monkeys[j].items = nil
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspects)))

	return inspects[0] * inspects[1]
}

func a(input string) {
	monkeys, _ := parse(input)
	fmt.Println(inspect(monkeys, 20, func(w int) int { return w / 3 }))
}

func b(input string) {
	monkeys, lcm := parse(input)
	fmt.Println(inspect(monkeys, 10000, func(w int) int { return w % lcm }))
}
