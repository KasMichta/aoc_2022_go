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

func checkVisible(trees map[[2]int]int, lim int) bool {

	return true
}

func sumVisible(trees map[[2]int]int, rowCount int, colCount int, s *int) {
	for cord, tree := range trees {
		row, col := cord[0], cord[1]
		if col != 0 && col != colCount-1 && row != 0 && row != rowCount-1 {
			l := true
			for i := col - 1; i >= 0; i-- {
				if trees[[2]int{row, i}] >= tree {
					l = false
				}
			}
			u := true
			for i := row - 1; i >= 0; i-- {
				if trees[[2]int{i, col}] >= tree {
					u = false
				}
			}

			r := true
			for i := col + 1; i <= colCount; i++ {
				if trees[[2]int{row, i}] >= tree {
					r = false
				}
			}

			d := true
			for i := row + 1; i <= rowCount; i++ {
				if trees[[2]int{i, col}] >= tree {
					d = false
				}
			}
			if l || r || u || d {
				*s++
			}
		}
	}
}

func main() {
	lines := readLines(os.Args[1])
	trees := make(map[[2]int]int)
	for i, line := range lines {
		row := strings.Split(line, "")
		for j, tree := range row {
			treeHgt, _ := strconv.Atoi(tree)
			cords := [2]int{i, j}
			trees[cords] = treeHgt
		}
	}
	fmt.Println(len(trees))
	var sum int
	sumVisible(trees, len(lines), len(lines[0]), &sum)
    edgeTrees := (len(lines)+len(lines[0]))*2 - 4
	fmt.Println(sum + edgeTrees)
}
