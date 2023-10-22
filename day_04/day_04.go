package main

import (
	"fmt"
	"os"
	"strconv"
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

func main() {
	//Part 1
	lines := readLines(os.Args[1])
	fmt.Printf("%v\n", lines[len(lines)-1:])

	var overlapTotal int
	for _, line := range lines {
		pair := strings.Split(line, ",")
		elfs := make(map[int][]int)

		for i, worker := range pair {
			workString := strings.Split(worker, "-")
			var work []int

			for _, w := range workString {
				workInt, err := strconv.Atoi(w)
				if err != nil {
					panic(err)
				}
				work = append(work, workInt)
			}
			elfs[i] = work
		}

		//one elfwork contains the other and vice versa
		if (elfs[0][0] <= elfs[1][0]) && (elfs[0][1] >= elfs[1][1]) {
			overlapTotal++
		} else if (elfs[1][0] <= elfs[0][0]) && (elfs[1][1] >= elfs[0][1]) {
			overlapTotal++
		}
	}
	fmt.Printf("total: %v\n", overlapTotal)

    //Part 2
}

