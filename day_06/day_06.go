package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func readFile(file string) []string {
	//load file
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(data), "")
	return content
}

func verifyMarker(curr []string) bool {
	for i, char := range curr {
		if slices.Contains(curr[i+1:], char) {
			return false
		}
	}
	return true
}

func searchMarker(strmBuff []string, length int) {
	for i := 0; i <= len(strmBuff)-length; i++ {
		curr := strmBuff[i : i+length]
		if verifyMarker(curr) {
			//last char of curr
			fmt.Printf("at: %v\n", i+length)
			fmt.Printf("marker: %v\n", curr)
			break
		}
	}
}

func main() {
	//Part 1
	dataStrmBuff := readFile(os.Args[1])
    searchMarker(dataStrmBuff, 4)
    //Part 2
    searchMarker(dataStrmBuff, 14)
}
