package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func checkErr(e error) bool {
	if e != nil {
		fmt.Println(e)
		return true
	} else {
		return false
	}
}

func main() {
	var inputString = "input.txt"
	var inputFilePath = filepath.FromSlash(inputString)
	fileContents, err := os.ReadFile(inputFilePath)

	if checkErr(err) {
		os.Exit(1)
	}

	indexVals := strings.Fields(string(fileContents))
	list1, list2, err := splitIndices(indexVals)
	if checkErr(err) {
		os.Exit(1)
	}
	diffs, err := subtractIndices(list1, list2)

	if checkErr(err) {
		os.Exit(1)
	}

	total := sumDiffs(diffs)

	fmt.Println("Total:", total)
}

func sumDiffs(diffs []int) int {
	var total = 0

	for _, diff := range diffs {
		total += diff
	}
	return total
}

func subtractIndices(l1 []int, l2 []int) ([]int, error) {
	if len(l1) != len(l2) {
		return nil, fmt.Errorf("lengths not equal")
	}

	slices.Sort(l1)
	slices.Sort(l2)

	diffs := make([]int, len(l1))
	for i := 0; i < len(l1); i += 1 {
		value1 := l1[i]
		value2 := l2[i]

		diff := value1 - value2
		posDiff := int(math.Abs(float64(diff)))
		diffs[i] = posDiff
	}
	return diffs, nil
}

func splitIndices(indexVals []string) ([]int, []int, error) {
	l1 := make([]int, len(indexVals)/2)
	l2 := make([]int, len(indexVals)/2)

	for i := 0; i < (len(indexVals) / 2); i += 1 {
		i1 := i * 2
		i2 := i1 + 1
		int1, err := strconv.Atoi(indexVals[i1])
		int2, err2 := strconv.Atoi(indexVals[i2])
		if checkErr(err) || checkErr(err2) {
			return nil, nil, fmt.Errorf("unable to convert indices")
		}

		l1[i] = int1
		l2[i] = int2
	}
	return l1, l2, nil
}
