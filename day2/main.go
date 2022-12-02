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
	b(scanner)
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
