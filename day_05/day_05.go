package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
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

		for i := 0; i < (breakLine - 1); i++ {
			crate := string(lines[i][column])
			if crate != " " {
				crates = append(crates, crate)
			}
		}
		stacks[stkid+1] = crates
	}
	return stacks
}

func readProcedure(procedureString string) (move, from, to int) {
	re := regexp.MustCompile(`\d+`)
	steps := re.FindAllString(procedureString, -1)

	move, _ = strconv.Atoi(steps[0])
	from, _ = strconv.Atoi(steps[1])
	to, _ = strconv.Atoi(steps[2])

	return
}

func main() {
	//Part 1
	lines := readLines(os.Args[1])
	procLines := lines[10:]
	//fmt.Printf("%v\n", readCrates(lines))
	//fmt.Printf("%v\n", readProcedure(lines[10]))
	stacksP1 := readCrates(lines)
	//Part 2
	stacksP2 := readCrates(lines)
	for _, line := range procLines {
		move, from, to := readProcedure(line)

		cratesToMoveP1 := stacksP1[from][:move]
		cratesToMoveP2 := stacksP2[from][:move]
		//because they are LIFO
		slices.Reverse(cratesToMoveP1)
		stacksP1[from] = stacksP1[from][move:]
		stacksP2[from] = stacksP2[from][move:]

		stacksP1[to] = slices.Insert(stacksP1[to], 0, cratesToMoveP1...)
		stacksP2[to] = slices.Insert(stacksP2[to], 0, cratesToMoveP2...)
	}

	stackKeys := maps.Keys(stacksP1)
	slices.Sort(stackKeys)
	var topOfStacksP1 string
	var topOfStacksP2 string
	for _, stack := range stackKeys {
		topOfStacksP1 += stacksP1[stack][0]
		topOfStacksP2 += stacksP2[stack][0]
	}
	fmt.Printf("Part 1: %v\nPart 2: %v\n", topOfStacksP1, topOfStacksP2)
}
