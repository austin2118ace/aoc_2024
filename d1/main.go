package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
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
	var inputString = "C:\\Projects\\aoc_2024\\d1\\input.txt"
	var inputFilePath = filepath.FromSlash(inputString)
	fileContents, err := os.ReadFile(inputFilePath)

	if checkErr(err) {
		os.Exit(1)
	}

	indexVals := strings.Split(string(fileContents), "   ")
	_, err = subtractIndices(indexVals)
	fmt.Println(err)
}

func prepIndices(indexVals []string) ([]string, error) {
	return nil, nil
}

func subtractIndices(indexVals []string) ([]int, error) {

	for i := 0; i < (len(indexVals) - 1); i += 2 {
		value1 := indexVals[i]
		value2 := indexVals[i+1]
		int1, err := strconv.Atoi(value1)
		int2, err2 := strconv.Atoi(value2)
		if checkErr(err) || checkErr(err2) {
			return nil, fmt.Errorf("unable to subtract indices")
		}

		diff := int1 - int2
		posDiff := math.Abs(float64(diff))
		fmt.Println(posDiff)
	}
	return nil, nil
}
