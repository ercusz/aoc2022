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

type Pair struct {
	l []any
	r []any
}

func parse(input string) ([]Pair, []any) {
	lines := strings.Split(input, "\n\n")
	pairs := make([]Pair, len(lines))
	packets := []any{}
	for idx, line := range lines {
		tmp := strings.Split(line, "\n")
		json.Unmarshal([]byte(tmp[0]), &pairs[idx].l)
		json.Unmarshal([]byte(tmp[1]), &pairs[idx].r)
		packets = append(packets, pairs[idx].l, pairs[idx].r)
	}

	return pairs, packets
}

func main() {
	a(input)
	b(input)
}

func compare(left, right any) int {
	l, okL := left.(float64)
	r, okR := right.(float64)
	if okL && okR {
		return int(l) - int(r)
	}

	var fList, sList []any

	switch left.(type) {
	case []any, []float64:
		fList = left.([]any)
	case float64:
		fList = []any{left}
	}

	switch right.(type) {
	case []any, []float64:
		sList = right.([]any)
	case float64:
		sList = []any{right}
	}

	for i := range fList {
		if len(sList) <= i {
			return 1
		}
		if r := compare(fList[i], sList[i]); r != 0 {
			return r
		}
	}

	if len(sList) == len(fList) {
		return 0
	}

	return -1
}

func a(input string) {
	pairs, _ := parse(input)

	sum := 0
	for idx, pair := range pairs {
		if compare(pair.l, pair.r) <= 0 {
			sum += idx + 1
		}
	}

	fmt.Println(sum)
}

func b(input string) {
	_, packets := parse(input)

	var pkt1, pkt2 any
	json.Unmarshal([]byte("[[2]]"), &pkt1)
	json.Unmarshal([]byte("[[6]]"), &pkt2)
	packets = append(packets, pkt1, pkt2)

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) < 0
	})

	r := 1
	for k, v := range packets {
		str, _ := json.Marshal(v)
		if string(str) == "[[2]]" || string(str) == "[[6]]" {
			r *= k + 1
		}
	}

	fmt.Println(r)
}
