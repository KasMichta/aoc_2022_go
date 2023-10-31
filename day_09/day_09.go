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

func isTouching(h []int, t []int) bool {
	if (h[0] == t[0] || h[0] == t[0]+1 || h[0] == t[0]-1) &&
		(h[1] == t[1] || h[1] == t[1]+1 || h[1] == t[1]-1) {
		return true
	} else {
		return false
	}
}

func moveTail(h []int, t []int) {
	if h[0] == t[0] {
		if h[1] == t[1]-2 {
			t[1]--
		} else if h[1] == t[1]+2 {
			t[1]++
		}
	} else if h[1] == t[1] {
		if h[0] == t[0]-2 {
			t[0]--
		} else if h[0] == t[0]+2 {
			t[0]++
		}
	} else {
		if h[0] == t[0]+2 {
			t[0]++
			if h[1] > t[1] {
				t[1]++
			} else {
				t[1]--
			}

		} else if h[0] == t[0]-2 {
			t[0]--
			if h[1] > t[1] {
				t[1]++
			} else {
				t[1]--
			}
		} else if h[1] == t[1]+2 {
			t[1]++
			if h[0] > t[0] {
				t[0]++
			} else {
				t[0]--
			}
		} else if h[1] == t[1]-2 {
			t[1]--
			if h[0] > t[0] {
				t[0]++
			} else {
				t[0]--
			}
		}
	}
}

func main() {
	//Part 1
	lines := readLines(os.Args[1])
	head := []int{0, 0}
	tail := []int{0, 0}
	tailPositions := make(map[string]bool)

	//Part 2
	rope := make([][]int, 10)
	for i := 0; i <= 9; i++ {
		rope[i] = []int{0, 0}
	}
	ropeTailPositions := make(map[string]bool)

	for _, line := range lines {
		op := strings.Split(line, " ")
		d := op[0]
		m, _ := strconv.Atoi(op[1])
		for i := 1; i <= m; i++ {
			switch d {
			case "L":
				head[0]--
				rope[0][0]--
			case "R":
				head[0]++
				rope[0][0]++
			case "U":
				head[1]--
				rope[0][1]--
			case "D":
				head[1]++
				rope[0][1]++
			}

			if !isTouching(head, tail) {
				moveTail(head, tail)
			}
			xy := fmt.Sprint(tail[0], tail[1])
			tailPositions[xy] = true

			endOfRope := len(rope) - 1
			for i := 0; i < endOfRope; i++ {
				if !isTouching(rope[i], rope[i+1]) {
					moveTail(rope[i], rope[i+1])
				}
				if i+1 == endOfRope {
					xy := fmt.Sprint(rope[endOfRope][0], rope[endOfRope][1])
					ropeTailPositions[xy] = true
				}
			}
		}
	}

	count := 0
	for _, v := range tailPositions {
		if v {
			count++
		}
	}

	ropeCount := 0
	for _, v := range ropeTailPositions {
		if v {
			ropeCount++
		}
	}
	fmt.Println("unique tail positions:", count)
	fmt.Println("unique rope tail positions:", ropeCount)
}
