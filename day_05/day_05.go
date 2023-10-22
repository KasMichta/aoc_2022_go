package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
    "strconv"
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
	stackLine := lines[breakLine-1]

	re := regexp.MustCompile(`\d+`)
	stackNumbers := re.FindAllStringIndex(stackLine, -1)

	stacks := make(map[int][]string)

	for stkid, stk := range stackNumbers {
        //index of stack column
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

func readProcedure(procedureString string) map[string]int {
    re:= regexp.MustCompile(`\d+`)
    steps := re.FindAllString(procedureString, -1)
    procedure := make(map[string]int)

    procedure["move"], _ = strconv.Atoi(steps[0])
    procedure["from"], _ = strconv.Atoi(steps[1])
    procedure["to"], _ = strconv.Atoi(steps[2])

    return procedure
}

func main() {
	lines := readLines(os.Args[1])
	fmt.Printf("%v\n", readCrates(lines))
    fmt.Printf("%v\n", readProcedure(lines[10]))
}
