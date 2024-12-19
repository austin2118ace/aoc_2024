package main

import (
	"fmt"
	"os"
	"path/filepath"
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

	indexVals := strings.Split(string(fileContents), " ")
	fmt.Println(indexVals)
}
