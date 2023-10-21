package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "slices"
)

//helper for calls
func check(e error) {
    if e != nil {
        panic (e)
    }
}

func sumTopX (slc []int, x int) int{
    sum := 0
    topX := slc[len(slc) - x:]
    for _, r := range topX {
        sum += r
    }
    return sum
}

func main() {
    data, err := os.ReadFile(os.Args[1])
    check(err)
    
    splitLines := strings.Split(string(data), "\n\n")
    //fmt.Printf("Third Batch:\n%v\n", splitLines[2])

    var snackTotals []int
    
    //for each snack bag
    for _, r := range splitLines {
        snackTotal := 0
        splitSnacks := strings.Split(r, "\n")
        
        //convert calories to int and sum
        for _, s := range splitSnacks {
            snackVal, err := strconv.Atoi(s)
            if err != nil {
                snackTotal += 0
            }
            snackTotal += snackVal
        }
        //add the sums to slice
        snackTotals = append(snackTotals, snackTotal)
    }

    slices.Sort(snackTotals)

    sumTop3 := sumTopX(snackTotals, 3)
    fmt.Printf("Sum top 3: %v\n", sumTop3)
    //max := snackTotals[0]
    //for _, t := range snackTotals {
    //    if t > max {
    //        max = t
    //    }
    //}

    //fmt.Printf("%v\n",max)
}
