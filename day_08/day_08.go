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

func walkLeft(trees map[[2]int]int, c int, r int, t int) (bool, int) {
	visible := true
	for i := c - 1; i >= 0; i-- {
		if trees[[2]int{r, i}] >= t {
			visible = false
		}
	}
	return visible, 0
}

func walkRight(trees map[[2]int]int, c int, r int, lim int, t int) (bool, int) {
	visible := true
	for i := c + 1; i <= lim; i++ {
		if trees[[2]int{r, i}] >= t {
			visible = false
		}
	}
	return visible, 0
}

func walkUp(trees map[[2]int]int, c int, r int, t int) (bool, int) {
	visible := true
	for i := r - 1; i >= 0; i-- {
		if trees[[2]int{i, c}] >= t {
			visible = false
		}
	}
	return visible, 0
}

func walkDown(trees map[[2]int]int, c int, r int, lim int, t int) (bool, int) {
	visible := true
	for i := r + 1; i <= lim; i++ {
		if trees[[2]int{i, c}] >= t {
			visible = false
		}
	}
	return visible, 0
}

func sumVisible(trees map[[2]int]int, rowCount int, colCount int, s *int) {
	for cord, tree := range trees {
		row, col := cord[0], cord[1]
		if col != 0 && col != colCount-1 && row != 0 && row != rowCount-1 {
			l, _ := walkLeft(trees, col, row, tree)
			r, _ := walkRight(trees, col, row, colCount, tree)
			u, _ := walkUp(trees, col, row, tree)
			d, _ := walkDown(trees, col, row, rowCount, tree)

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

	var sum int
	sumVisible(trees, len(lines), len(lines[0]), &sum)
	edgeTrees := (len(lines)+len(lines[0]))*2 - 4
	fmt.Println(sum + edgeTrees)
}
