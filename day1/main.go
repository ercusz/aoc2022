package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Read input from file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	defer file.Close()

	// Create new scanner
	scanner := bufio.NewScanner(file)

	//part 1
	a(scanner)
	// part 2
	b(scanner)
}

func a(scanner *bufio.Scanner) {
	max := 0
	sum := 0 

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if (sum > max) {
				max = sum
			}
			sum = 0

			continue
		}

		curr, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		sum += curr
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result=%v", max)
}

func b(scanner *bufio.Scanner) {
	var sums []int
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			sums = append(sums, sum)
			sum = 0
			continue
		}

		curr, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		sum += curr
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	result := 0
	for i:=0; i<3; i++ {
		result += sums[i]
	}
	fmt.Printf("result=%v", result)
}