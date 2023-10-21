package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//Part 1
	data, err := os.ReadFile(os.Args[1])
	check(err)

	splitLines := strings.Split(string(data), "\n")
	//omit the last as it will be an empty string
	splitLines = splitLines[:len(splitLines)-1]

	strats := map[string]int{
		"A Y": 6,
		"B Z": 6,
		"C X": 6,
		"A X": 3,
		"B Y": 3,
		"C Z": 3,
		"A Z": 0,
		"B X": 0,
		"C Y": 0,
	}

	shapeScores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	var turnScores []int

	for _, turn := range splitLines {
		turnScore := shapeScores[string(turn[2])]
		turnScore += strats[string(turn)]

		turnScores = append(turnScores, turnScore)
	}

	var totalScore int
	for _, score := range turnScores {
		totalScore += score
	}

	fmt.Printf("Part 1: %v\n", totalScore)

	//Part 2
	shapeScores = map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

    outcomeScores := map[string]int{
        "X" : 0,
        "Y" : 3,
        "Z" : 6,
    }

	shapeStrats := map[string]map[string]string{
		"A": {
			"X": "Scissors",
			"Y": "Rock",
			"Z": "Paper",
		},
		"B": {
			"X": "Rock",
			"Y": "Paper",
			"Z": "Scissors",
		},
		"C": {
			"X": "Paper",
			"Y": "Scissors",
			"Z": "Rock",
		},
	}

    var roundScores int

    for _, round := range splitLines {
        thrownShape := string(round[0])
        outcome := string(round[2])
        //score for choosing the appropriate response
        roundScore := shapeScores[shapeStrats[thrownShape][outcome]]
        //score for outcome
        roundScore += outcomeScores[outcome]        
        roundScores += roundScore
    }

    fmt.Printf("Part 2: %v\n", roundScores)

}
