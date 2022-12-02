package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	// a(scanner)
	// b(scanner)
	// a_matrix(scanner)
	b_matrix(scanner)
}

var points = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func a(scanner *bufio.Scanner) {
	myScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		choice := strings.Split(line, " ")
		// A X = Rock
		// B Y = Paper
		// C Z = Scissor
		// x=1, y=2, z=3
		// A win Z | A lose Y
		// B win X | B lose C
		// C win Y | C lose X
		elvesChoice := choice[0]
		myChoice := choice[1]
		switch myChoice {
		case "X":
			if elvesChoice == "C" {
				myScore += points[myChoice] + 6
				fmt.Printf("You Win | %v\n", choice)
				continue
			} else if elvesChoice == "B" {
				myScore += points[myChoice] + 0
				fmt.Printf("You Lose | %v\n", choice)
				continue
			}

			myScore += points[myChoice] + 3
			fmt.Printf("You Draw | %v\n", choice)
			continue
		case "Y":
			if elvesChoice == "A" {
				myScore += points[myChoice] + 6
				fmt.Printf("You Win | %v\n", choice)
				continue
			} else if elvesChoice == "C" {
				myScore += points[myChoice] + 0
				fmt.Printf("You Lose | %v\n", choice)
				continue
			}

			myScore += points[myChoice] + 3
			fmt.Printf("You Draw | %v\n", choice)
			continue
		case "Z":
			if elvesChoice == "B" {
				myScore += points[myChoice] + 6
				fmt.Printf("You Win | %v\n", choice)
				continue
			} else if elvesChoice == "A" {
				myScore += points[myChoice] + 0
				fmt.Printf("You Lose | %v\n", choice)
				continue
			}

			myScore += points[myChoice] + 3
			fmt.Printf("You Draw | %v\n", choice)
			continue
		}
	}

	fmt.Println(myScore)
}

func b(scanner *bufio.Scanner) {
	var endResult = map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	var choicePoint = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	var choicePoint2 = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	myScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		choice := strings.Split(line, " ")
		er := endResult[choice[1]]
		if er == 0 {
			// lose
			if choice[0] == "A" {
				myScore += er + choicePoint["Z"]
			} else if choice[0] == "B" {
				myScore += er + choicePoint["X"]
			} else if choice[0] == "C" {
				myScore += er + choicePoint["Y"]
			}
			continue
		} else if er == 3 {
			// draw
			myScore += er + choicePoint2[choice[0]]
			continue
		} else if er == 6 {
			// win
			if choice[0] == "A" {
				myScore += er + choicePoint["Y"]
			} else if choice[0] == "B" {
				myScore += er + choicePoint["Z"]
			} else if choice[0] == "C" {
				myScore += er + choicePoint["X"]
			}
			continue
		}
	}

	fmt.Println(myScore)
}

func a_matrix(scanner *bufio.Scanner) {
	// Your puzzle answer was 10310.
	// +--------------+----------+-----------+--------------+----+
	// |              | ROCK (X) | PAPER (Y) | SCISSORS (Z) | ME |
	// +--------------+----------+-----------+--------------+----+
	// | (A) ROCK     | DRAW     | ME        | ELVES        |    |
	// +--------------+----------+-----------+--------------+----+
	// | (B) PAPER    | ELVES    | DRAW      | ME           |    |
	// +--------------+----------+-----------+--------------+----+
	// | (C) SCISSORS | ME       | ELVES     | DRAW         |    |
	// +--------------+----------+-----------+--------------+----+
	// | ELVES        |          |           |              |    |
	// +--------------+----------+-----------+--------------+----+
	myChoices := map[string]int{
		"X": 0,
		"Y": 1,
		"Z": 2,
	}

	elvesChoices := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
	}
	compareTable := [][]int{
		{3 + 1, 6 + 2, 0 + 3},
		{0 + 1, 3 + 2, 6 + 3},
		{6 + 1, 0 + 2, 3 + 3},
	}
	myScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		choice := strings.Split(line, " ")
		myScore += compareTable[elvesChoices[choice[0]]][myChoices[choice[1]]]
	}
	fmt.Println(myScore)
}

func b_matrix(scanner *bufio.Scanner) {
	// Your puzzle answer was 14859.
	// +--------------+------------------+------------------+------------------+
	// |              | LOSE (X)         | DRAW (Y)         | WIN (Z)          |
	// +--------------+------------------+------------------+------------------+
	// | (A) ROCK     | 0 + 3 (SCISSORS) | 3 + 1 (ROCK)     | 6 + 2 (PAPER)    |
	// +--------------+------------------+------------------+------------------+
	// | (B) PAPER    | 0 + 1 (ROCK)     | 3 + 2 (PAPER)    | 6 + 3 (SCISSORS) |
	// +--------------+------------------+------------------+------------------+
	// | (C) SCISSORS | 0 + 2 (PAPER)    | 3 + 3 (SCISSORS) | 6 + 1 (ROCK)     |
	// +--------------+------------------+------------------+------------------+
	// | ELVES        |                  |                  |                  |
	// +--------------+------------------+------------------+------------------+
	resultChoices := map[string]int{
		"X": 0,
		"Y": 1,
		"Z": 2,
	}

	elvesChoices := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
	}
	compareTable := [][]int{
		{0 + 3, 3 + 1, 6 + 2},
		{0 + 1, 3 + 2, 6 + 3},
		{0 + 2, 3 + 3, 6 + 1},
	}
	myScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		choice := strings.Split(line, " ")
		myScore += compareTable[elvesChoices[choice[0]]][resultChoices[choice[1]]]
	}
	fmt.Println(myScore)
}
