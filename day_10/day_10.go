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

type clock struct {
	cycle  int
	reg    map[string]int
	buffer map[string][][]int
}

func (c *clock) tick() {
	if c.buffer != nil {
		for k, b := range c.buffer {
			for _, op := range b {
				if op[0] == 1 {
					c.reg[k] += op[1]
                    c.buffer[k] = c.buffer[k][1:]
				} else {
					op[0]--
				}
			}
		}
	}

	c.cycle++
}

func (c *clock) noop() {
	c.tick()
}

func (c *clock) add(reg string, val int) {
	op := []int{2, val}
	c.buffer[reg] = append(c.buffer[reg], op)
}

func main() {
	lines := readLines(os.Args[1])
	reg := map[string]int{"x": 1}

	buffer := map[string][][]int{
		"x": make([][]int, 0),
	}

	clock := clock{
		cycle:  1,
		reg:    reg,
		buffer: buffer,
	}

    var remainingCycles int

	for _, line := range lines {
		fmt.Println("=========")
		fmt.Println(clock.cycle, clock.reg, clock.buffer)
		fmt.Println(line)

		if line == "noop" {
            remainingCycles++
			clock.noop()
		} else {
            remainingCycles += 2
			parts := strings.Split(line, " ")
			reg := string(parts[0][len(parts[0])-1])
			val, _ := strconv.Atoi(parts[1])
			clock.add(reg, val)
            clock.tick()
		}
		fmt.Println(clock.cycle, clock.reg, clock.buffer)
        remainingCycles--
	}

    for remainingCycles > 0 {
        clock.tick()
        remainingCycles--
        fmt.Println(clock.cycle, clock.reg, clock.buffer)
    }
}
