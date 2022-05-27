package main

import (
	"fmt"
	"os"
)

func main() {
	intervalOne, _ := GenerateIntervalWithValue(1, 10, "a")
	intervalTwo, _ := GenerateIntervalWithValue(11, 20, "b")
	intervalThree, _ := GenerateIntervalWithValue(21, 30, "c")
	intervalFour, _ := GenerateIntervalWithValue(31, 40, "d")
	intervalFive, _ := GenerateIntervalWithValue(41, 50, "e")
	intervals := []IntervalWithValue[int]{intervalOne, intervalFive, intervalFour, intervalTwo, intervalThree}
	tree, _ := GenerateTree(intervals, false)
	result, err := tree.SearchFirstInterval(41, true, false)
	fmt.Fprintln(os.Stdout, err, result)
}
