package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	//"slices"
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

func countVisible(trees []int, t *int) {
	innerTrees := trees[1 : len(trees)-1]
	//fmt.Println(len(innerTrees), cap(innerTrees))
	//firstTree := trees[0]
	tallestTree := trees[0]
	for _, tree := range innerTrees {
		if tree > tallestTree {
			tallestTree = tree
			*t++
		}
	}
}

func main() {
	lines := readLines(os.Args[1])
	rowCount := len(lines)
	colCount := len(lines[0])
	rows := make([][]int, rowCount)
	cols := make([][]int, colCount)
	trees := map[string][][]int{
		"rows": rows,
		"cols": cols,
	}
	for i, line := range lines {
		rowStrs := strings.Split(line, "")
		for j, tree := range rowStrs {
			treeHgt, _ := strconv.Atoi(tree)
			rows[i] = append(rows[i], treeHgt)
			cols[j] = append(cols[j], treeHgt)
		}
	}
	var firstTotal int
	fmt.Println(trees["rows"][20])
	countVisible(trees["rows"][20], &firstTotal)
	fmt.Println(firstTotal)

	var totalVisible int
	for _, row := range trees["rows"] {
		//L->R
		countVisible(row, &totalVisible)
		//R->L
		slices.Reverse(row)
		countVisible(row, &totalVisible)
	}
	for _, col := range trees["cols"] {
		//T->B
		countVisible(col, &totalVisible)
		//B->T
		slices.Reverse(col)
		countVisible(col, &totalVisible)
	}
	fmt.Println("innerTrees:", totalVisible)
	//fmt.Println("lines", len(lines))
	//fmt.Println("cols", len(lines[0]))
	//// remove 4 counted twice
	edgeTrees := (len(lines)+len(lines[0]))*2 - 4
	fmt.Println("totalTrees:", edgeTrees+totalVisible)
	fmt.Println(len(trees["rows"]))
    fmt.Println(len(trees["cols"]))
}
