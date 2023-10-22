package main

import (
	"os"
	"slices"
	"strings"
    "fmt"
)

func readLines(file string) []string {
	//load file
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	//split lines
	splitLines := strings.Split(string(data), "\n")
	splitLines = splitLines[:len(splitLines)-1]
	return splitLines
}

func readCrates(lines []string) map[int]string {
	breakLine := slices.Index(lines, "")
	fmt.Printf("break line: %v\n", breakLine)
	crates := make(map[int]string)
	return crates
}

func main() {
	lines := readLines(os.Args[1])
	readCrates(lines)
}
