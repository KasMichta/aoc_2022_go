package main

import (
	"os"
	"slices"
	"strings"
    "fmt"
    "regexp"
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
    
    crateLine := lines[breakLine-1]
    re := regexp.MustCompile(`\d+`)
    crateNumbers := re.FindAllString(crateLine, -1)
    
    fmt.Printf("crate numbers: %v\n", crateNumbers)
	crates := make(map[int]string)
	return crates
}

func main() {
	lines := readLines(os.Args[1])
	readCrates(lines)
}
