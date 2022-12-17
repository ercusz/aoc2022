package main

import (
	_ "embed"
	"encoding/json"
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

type Valve struct {
	rate      int
	leadTo    []string
	distances map[string]int
}

type State struct {
	valveName string
	visited   string
	pressure  int
	time      int
}

func parse(input string) map[string]*Valve {
	lines := strings.Split(input, "\n")
	valves := make(map[string]*Valve, len(lines))
	for _, line := range lines {
		var (
			valName, leadTo string
			v               Valve
		)

		line = strings.NewReplacer(
			" tunnel leads to valve ", "",
			" tunnels lead to valves ", "",
			", ", "\",\"",
		).Replace(line)

		fmt.Sscanf(
			line,
			"Valve %s has flow rate=%d;%s",
			&valName, &v.rate, &leadTo,
		)
		leadTo = "[\"" + leadTo + "\"]"
		json.Unmarshal([]byte(leadTo), &v.leadTo)
		v.distances = make(map[string]int)
		valves[valName] = &v
	}

	// find distances for each valve
	for start, valve := range valves {
		valve.distances[start] = 0

		queue := make(utils.Queue[string], 0)
		visited := make(map[string]bool)

		queue.Enqueue(start)
		for len(queue) > 0 {
			curr, _ := queue.Dequeue()
			if visited[curr] {
				continue
			}
			visited[curr] = true

			for _, child := range valves[curr].leadTo {
				if visited[child] {
					continue
				}
				valves[child].distances[start] = valves[curr].distances[start] + 1
				queue.Enqueue(child)
			}
		}
	}
	return valves
}

func main() {
	a(input)
	b(input)
}

func traversal(valves map[string]*Valve, time int) ([]string, int) {
	initial := State{
		valveName: "AA",
		visited:   "AA",
		pressure:  0,
		time:      time,
	}
	queue := make(utils.Queue[State], 0)
	pressures := make(map[string]int)

	queue.Enqueue(initial)
	for len(queue) > 0 {
		curr, _ := queue.Dequeue()
		valve := valves[curr.valveName]

		// add current visited path to pressures map
		path := curr.visited
		if _, ok := pressures[path]; !ok {
			pressures[path] = curr.pressure
		}

		// find reachable valves for the remaining time
		reachable := make([]string, 0)
		for next, distance := range valve.distances {
			if distance < curr.time {
				if !strings.Contains(curr.visited, next) {
					if valves[next].rate != 0 {
						reachable = append(reachable, next)
					}
				}
			}
		}

		// add next states to queue
		for _, finish := range reachable {
			newTime := curr.time - valve.distances[finish] - 1
			newPressure := newTime * valves[finish].rate

			next := State{
				valveName: finish,
				visited:   fmt.Sprintf("%s-%s", curr.visited, finish),
				pressure:  curr.pressure + newPressure,
				time:      newTime,
			}

			queue = append(queue, next)
		}
	}

	// find the max pressure and path of the max pressure
	var (
		maxPath     string
		maxPressure int
	)
	for path, pressure := range pressures {
		if pressure > maxPressure {
			maxPath = path
			maxPressure = pressure
		}
	}

	return strings.Split(maxPath, "-"), maxPressure
}

func a(input string) {
	valves := parse(input)
	_, pressure := traversal(valves, 30)
	fmt.Println(pressure)
}

func b(input string) {
	valves := parse(input)

	// I traverses
	path, pressure := traversal(valves, 26)

	// set the rate of the valves in the max path after I traversed to 0
	// to prevent the elephant from traversing the same path
	for _, name := range path {
		if v, ok := valves[name]; ok {
			v.rate = 0
			valves[name] = v
		}
	}

	// elephant traverses
	_, elephantPressure := traversal(valves, 26)

	fmt.Println(pressure + elephantPressure)
}
