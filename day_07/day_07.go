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

type fileSystem struct {
	rootDir directory
}

type directory struct {
	name      string
	parent    *directory
	children  []*directory
	totalSize int
}

func sumDirsOfSize(d *directory, size int, sum *int) {
	if d.children == nil {
		if d.totalSize <= size {
			*sum += d.totalSize
		}
	} else {
		for _, dir := range d.children {
			sumDirsOfSize(dir, size, sum)
			d.totalSize += dir.totalSize
		}
		if d.totalSize <= size {
			*sum += d.totalSize
		}
	}
}

func (d *directory) addDir(n string) {
	d.children = append(d.children, &directory{
		name:   n,
		parent: d,
	})
}

func (d *directory) addFile(l string) {
	fileSize := strings.Split(l, " ")
	size, _ := strconv.Atoi(fileSize[0])
	d.totalSize += size
}

func main() {
	lines := readLines(os.Args[1]) //[:36]
	fs := fileSystem{directory{name: "root"}}
	currDir := &fs.rootDir
	var i int
	for i < len(lines) {
		currLine := lines[i]
		switch {
		case strings.Contains(currLine, "$ ls"):
			i++
		case strings.Contains(currLine, "$ cd"):
			dirName := strings.Split(currLine, " ")[2]
			if dirName == ".." {
				currDir = currDir.parent
			} else {
				currDir.addDir(dirName)
				currDir = currDir.children[len(currDir.children)-1]
			}
			i++
		default:
			if !strings.Contains(currLine, "dir ") {
				currDir.addFile(currLine)
			}
			i++
		}
	}
	total := 0
	sumDirsOfSize(&fs.rootDir, 100000, &total)
	fmt.Println(total)
}
