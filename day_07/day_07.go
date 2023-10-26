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

func sumDirs(d *directory, total *int, thrsh int) {
	if d.children == nil {
		if d.totalSize <= thrsh {
			*total += d.totalSize
            fmt.Println(d.name, d.totalSize)
		}
	} else {
		for _, dir := range d.children {
            fmt.Println("inner", dir.name, dir.totalSize)
			sumDirs(dir, total, thrsh)
            d.totalSize += dir.totalSize
            fmt.Println("after children:",d.totalSize)
		}
		if d.totalSize <= thrsh {
			*total += d.totalSize
            fmt.Println(d.name, d.totalSize)
		}
	}
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
			//fmt.Println("ls")
			i++
		case strings.Contains(currLine, "$ cd"):
			dirName := strings.Split(currLine, " ")[2]
			//fmt.Println(dirName)
			if dirName == ".." {
				currDir = currDir.parent
				//fmt.Println("OutTo:", currDir.name)
			} else {
				currDir.children = append(currDir.children, &directory{
					name:   dirName,
					parent: currDir,
				})
				currDir = currDir.children[len(currDir.children)-1]
				//fmt.Println("InTo:", currDir.name)
			}
			i++
		default:
			if !strings.Contains(currLine, "dir ") {
				fileSize := strings.Split(currLine, " ")
				size, _ := strconv.Atoi(fileSize[0])
				currDir.totalSize += size
				//dir := *currDir
				//fmt.Printf("current:%v size:%v\n", dir.name, dir.totalSize)
			}
			i++
		}
	}
	fmt.Println(fs.rootDir)
	total := 0
	sumDirs(&fs.rootDir, &total, 100000)
	fmt.Println(total)
	//rootDir := fs.rootDir.children[0]
	//fmt.Println(*rootDir)

}
