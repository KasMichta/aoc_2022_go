package main

import (
	"fmt"
	"os"
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

func scopeTrees(trees [][]int) (int, int) {
    var visSum int
    var maxScnScr int
	for y, row := range trees {
		for x, tree := range row {
			if y == 0 || x == 0 || y == len(trees)-1 || x == len(row)-1 {
                visSum++
			} else {
                visL := true
                var scnScrL int
				for nxt := x - 1; nxt >= 0; nxt-- {
                    if trees[y][nxt] >= tree {
                        visL = false
                        scnScrL++
                        break
                    }
                    scnScrL++
				}

                visR := true
                var scnScrR int
				for nxt := x + 1; nxt < len(row); nxt++ {
                    if trees[y][nxt] >= tree {
                        visR = false
                        scnScrR++
                        break
                    }
                    scnScrR++
				}

                visU := true
                var scnScrU int
				for nxt := y - 1; nxt >= 0; nxt-- {
                    if trees[nxt][x] >= tree {
                        visU = false
                        scnScrU ++
                        break
                    }
                    scnScrU++
				}

                visD := true
                var scnScrD int
				for nxt := y + 1; nxt < len(trees); nxt++ {
                    if trees[nxt][x] >= tree {
                        visD = false
                        scnScrD ++
                        break
                    }
                    scnScrD++
				}

                if visL || visR || visU || visD { 
                    visSum++
                }
                scnScr := scnScrL * scnScrR * scnScrU * scnScrD
                if scnScr > maxScnScr {
                    maxScnScr = scnScr
                }
			}

		}
	}
	return visSum, maxScnScr
}

func main() {
	lines := readLines(os.Args[1])
	trees := make([][]int, len(lines))
	for y, line := range lines {
		treeRow := strings.Split(line, "")
		for _, tree := range treeRow {
			treeHgt, _ := strconv.Atoi(tree)
			trees[y] = append(trees[y], treeHgt)
		}
	}

    sumVisible, maxScnScr := scopeTrees(trees)
	fmt.Println("Sum of visible trees:", sumVisible, "|", "Max scenic score:", maxScnScr)

}
