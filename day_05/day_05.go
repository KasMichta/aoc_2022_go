package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
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

func readCrates(lines []string) map[int][]string {
	breakLine := slices.Index(lines, "")
	//fmt.Printf("break line: %v\n", breakLine)

	stackLine := lines[breakLine-1]
	re := regexp.MustCompile(`\d+`)
	stackNumbers := re.FindAllStringIndex(stackLine, -1)
	//fmt.Printf("stack numbers: %v\n", stackNumbers)
	stacks := make(map[int][]string)

	for stkid, stk := range stackNumbers {
		column := stk[0]
        var crates []string

        for i := 0; i < (breakLine - 1) ; i++  {
		    crate := string(lines[i][column])
            if crate != " " {
                crates = append(crates, crate)
            }
		}
        stacks[stkid+1] = crates
	}
	return stacks
}

func main() {
	lines := readLines(os.Args[1])
	fmt.Printf("%v\n", readCrates(lines))
}
