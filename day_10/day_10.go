package main

import (
	"fmt"
	"os"
	"sort"
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
	cycle          int
	reg            map[string]int
	signalInterval int
	signals        map[int]int
}

func (c *clock) tick() {
	for r := range c.reg {
		if c.cycle%c.signalInterval == 0 {
			c.signals[c.cycle] = c.reg[r] * c.cycle
		}
	}
	c.cycle++
}

func (c *clock) add(reg string, val int) {
	c.tick()
	c.tick()
	c.reg[reg] += val
}

func main() {
	lines := readLines(os.Args[1])
	reg := map[string]int{"x": 1}
	signals := make(map[int]int)

	clock := clock{
		cycle:          1,
		reg:            reg,
		signalInterval: 20,
		signals:        signals,
	}

	for _, line := range lines {
		if line == "noop" {
			clock.tick()
		} else {
			parts := strings.Split(line, " ")
			reg := parts[0][len(parts[0])-1]
			val, _ := strconv.Atoi(parts[1])
			clock.add(string(reg), val)
		}
	}
	keys := make([]int, 0)
	for k := range clock.signals {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var sumSignals int
	var signalOffset int
	for _, k := range keys {
		if k <= 220 && k%(20+signalOffset) == 0 {
			sumSignals += clock.signals[k]
			signalOffset += 40
		}
	}
	fmt.Println(sumSignals)
}
