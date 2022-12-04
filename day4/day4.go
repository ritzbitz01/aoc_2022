package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("day4_input.txt")

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

	overlapPairsFull := 0
	overlapPairs := 0
	
	var TEST_LIST []string
	TEST_LIST = append(TEST_LIST, "5-8,7-65")

	for _, line := range fileLines {
		assignments := strings.Split(line, ",")
		firstSet := assignments[0]
		secondSet := assignments[1]

		firstRange := strings.Split(firstSet, "-")
		secondRange := strings.Split(secondSet, "-")

		firstFirst, _ := strconv.Atoi(firstRange[0])
		firstSecond, _ := strconv.Atoi(firstRange[1])
		secondFirst, _ := strconv.Atoi(secondRange[0])
		secondSecond, _ := strconv.Atoi(secondRange[1])

		if (((firstFirst >= secondFirst) && (firstSecond <= secondSecond)) ||
			((secondFirst >= firstFirst) && (secondSecond <= firstSecond))) {
				overlapPairsFull ++
			}

		if (((firstFirst >= secondFirst) && (firstSecond <= secondSecond)) ||
			((secondFirst >= firstFirst) && (secondSecond <= firstSecond)) ||
			((firstFirst <= secondFirst) && (firstSecond >= secondFirst)) ||
			((secondFirst <= firstFirst) && (secondSecond >= firstFirst))) {
				overlapPairs ++
			}
	}

	fmt.Println("OVERLAP FULL PAIRS: ", overlapPairsFull);
	fmt.Println("OVERLAP PAIRS: ", overlapPairs);

}
