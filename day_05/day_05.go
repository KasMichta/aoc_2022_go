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

func readProcedure(procedureString string) map[string]int {
	re := regexp.MustCompile(`\d+`)
	steps := re.FindAllString(procedureString, -1)
	procedure := make(map[string]int)

	procedure["move"], _ = strconv.Atoi(steps[0])
	procedure["from"], _ = strconv.Atoi(steps[1])
	procedure["to"], _ = strconv.Atoi(steps[2])

	return procedure
}

//func execProcedure(proc map[string]int, stacks map[int][]string) {
//	stkFrom := stacks[proc["from"]]
//	stkTo := stacks[proc["to"]]
//	var toMove []string
//	if len(stkFrom) >= proc["move"] {
//		toMove = stkFrom[:proc["move"]]
//		stkFrom = stkFrom[proc["move"]:]
//	} else {
//		toMove = stkFrom[:]
//		stkFrom = stkFrom[:]
//	}
//	targetStack := stkTo
//	//... unpacks b,
//	targetStack = append(targetStack, toMove...)
//}

//func execProcedure(move int, from int, to int, stackFrm *[]string, stackTo *[]string) {
//	var toMove []string
//	if len(&stackFrm) >= move {
//		toMove = &stackFrm[:move]
//		*stackFrm = &stackFrm[move:]
//	} else {
//		toMove = stackFrm[:]
//		stackFrm = stackFrm[:]
//	}
//	stackTo = append(stackTo, toMove...)
//}

func main() {
    //Part 1
	lines := readLines(os.Args[1])
	procLines := lines[10:]
	//fmt.Printf("%v\n", readCrates(lines))
	//fmt.Printf("%v\n", readProcedure(lines[10]))
	stacks := readCrates(lines)
	for _, line := range procLines {
		pr := readProcedure(line)
		stackFrom := pr["from"]
		stackTo := pr["to"]
		move := pr["move"]

		cratesToMove := stacks[stackFrom][:move]
        //because they are LIFO
        slices.Reverse(cratesToMove)
		stacks[stackFrom] = stacks[stackFrom][move:]

		stacks[stackTo] = slices.Insert(stacks[stackTo], 0, cratesToMove...)

	}

	stackKeys := maps.Keys(stacks)
    slices.Sort(stackKeys)
	for _, stack := range stackKeys {
		fmt.Printf("%v", stacks[stack][0])
	}
	fmt.Println()
}
