package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//Part 1
	//init priorities
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	priorities := make(map[string]int)

	prioInt := 1
	for _, letter := range letters {
		priorities[string(letter)] = prioInt
		prioInt++
	}

	fmt.Printf("priorities:\n%v\n", priorities)

	//load file
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	splitLines := strings.Split(string(data), "\n")
	//omit the last as it will be an empty string
	splitLines = splitLines[:len(splitLines)-1]

	//splitLines = splitLines[:5]

	var prioSum int
	for _, line := range splitLines {
		comp1 := line[:(len(line) / 2)]
		comp2 := line[(len(line) / 2):]

		for _, char := range comp1 {
			if strings.ContainsRune(comp2, char) {
                //remove char from both avoid multihits
                comp1 = strings.ReplaceAll(comp1, string(char), "")
                comp2 = strings.ReplaceAll(comp2, string(char), "")

				prioSum += priorities[string(char)]
			}
		}
	}
    fmt.Printf("Sum of Priorities: %v\n", prioSum)
    
}
