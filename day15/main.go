package main

import (
	_ "embed"
	"fmt"
	"image"
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
	a(input)
	b(input)
}

type sensor struct {
	image.Point
	beaconX, beaconY,
	distanceToClosestBeacon int
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return utils.AbsDiffInt(x1, x2) + utils.AbsDiffInt(y1, y2)
}

func parse(input string) (sensors []sensor) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var s sensor
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.X, &s.Y, &s.beaconX, &s.beaconY)
		s.distanceToClosestBeacon = manhattanDistance(s.X, s.Y, s.beaconX, s.beaconY)
		sensors = append(sensors, s)
	}
	return
}

func a(input string) {
	sensors := parse(input)
	y := 2000000
	currPoints := make(map[int]struct{})

	for _, sensor := range sensors {
		distanceFromY := utils.AbsDiffInt(sensor.Y, y)

		if distanceFromY > sensor.distanceToClosestBeacon {
			continue
		}

		dy := sensor.distanceToClosestBeacon - distanceFromY
		minIndex, maxIndex := sensor.X-dy, sensor.X+dy

		for x := minIndex; x <= maxIndex; x++ {
			if x == sensor.beaconX && y == sensor.beaconY {
				continue
			}

			currPoints[x] = struct{}{}
		}
	}

	fmt.Println(len(currPoints))
}

func b(input string) {
	sensors := parse(input)
	row := []image.Point{}
	nRows := 4000000

	for y := 0; y < nRows; y++ {
		for _, sensor := range sensors {
			distanceFromY := utils.AbsDiffInt(sensor.Y, y)

			if distanceFromY > sensor.distanceToClosestBeacon {
				continue
			}

			dy := sensor.distanceToClosestBeacon - distanceFromY
			minIndex, maxIndex := utils.MaxInt(0, sensor.X-dy), utils.MinInt(nRows, sensor.X+dy)
			row = append(row, image.Point{minIndex, maxIndex})
		}

		sort.Slice(row, func(i, j int) bool {
			return row[i].X < row[j].X
		})

		for j := 1; j < len(row); j++ {
			// replaced invalid row
			if row[j].Y < row[j-1].Y {
				row[j] = row[j-1]
			}
			// find the gap
			if row[j].X-row[j-1].Y == 2 {
				fmt.Println((uint64(row[j].X-1) * 4000000) + uint64(y))
				return
			}
		}

		row = nil
	}

}
