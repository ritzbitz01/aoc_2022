package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findMarker(part int) int {
	readFile, err := os.Open("day8_input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var charArray [99][99]int
	rowNum := 0
	for _, line := range fileLines {
		for idx, e := range line {
			charArray[rowNum][idx], _ = strconv.Atoi(string(e))
		}
		rowNum++
	}
	fmt.Println(charArray[98][98])
	return -1
}

func main() {
	fmt.Println("PART 1 INDEX: ", findMarker(1))
	fmt.Println("PART 2 INDEX: ", findMarker(2))
}
