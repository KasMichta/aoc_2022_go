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

func propDirSize(d *directory) {
	if d.children != nil {
		for _, dir := range d.children {
			propDirSize(dir)
			d.totalSize += dir.totalSize
		}
	}
}

func sumDirsOfSize(d *directory, size int, sum *int) {
	for _, dir := range d.children {
		sumDirsOfSize(dir, size, sum)
	}
	if d.totalSize <= size {
		*sum += d.totalSize
	}
}

func findDirSize(d *directory, reqSpace int, minSize *int) {
	for _, dir := range d.children {
		findDirSize(dir, reqSpace, minSize)
	}
	if d.totalSize >= reqSpace && d.totalSize < *minSize {
		*minSize = d.totalSize
	}
}

func main() {
	//Part 1
	lines := readLines(os.Args[1])
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

	propDirSize(&fs.rootDir)

	var total int
	threshold := 100000
	sumDirsOfSize(&fs.rootDir, threshold, &total)
	fmt.Printf("total of dirs under threshold: %v\n", total)

	//Part 2
	diskSize := 70000000
	updateSize := 30000000
	freeSpace := diskSize - fs.rootDir.totalSize
	reqSpace := updateSize - freeSpace
	minSize := freeSpace
	findDirSize(&fs.rootDir, reqSpace, &minSize)
	fmt.Printf("min dir size to free: %v\n", minSize)
}
