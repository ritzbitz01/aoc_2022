package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, err := os.Open("day2_input.txt")

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

	rpsOutcomes1 := make(map[string]int)
	rpsOutcomes1["A X"] = 4
	rpsOutcomes1["A Y"] = 8
	rpsOutcomes1["A Z"] = 3
	rpsOutcomes1["B X"] = 1
	rpsOutcomes1["B Y"] = 5
	rpsOutcomes1["B Z"] = 9
	rpsOutcomes1["C X"] = 7
	rpsOutcomes1["C Y"] = 2
	rpsOutcomes1["C Z"] = 6
	var outcomeSum = 0
	for _, line := range fileLines {
		fmt.Println(line)
		if len(line) > 0 {
			outcomeSum += rpsOutcomes1[line]
		} else {
			fmt.Println("ALL DONE!")
		}
	}

	rpsOutcomes2 := make(map[string]int)
	rpsOutcomes2["A X"] = 3
	rpsOutcomes2["A Y"] = 4
	rpsOutcomes2["A Z"] = 8
	rpsOutcomes2["B X"] = 1
	rpsOutcomes2["B Y"] = 5
	rpsOutcomes2["B Z"] = 9
	rpsOutcomes2["C X"] = 2
	rpsOutcomes2["C Y"] = 6
	rpsOutcomes2["C Z"] = 7

	var outcomeSum2 = 0
	for _, line := range fileLines {
		fmt.Println(line)
		if len(line) > 0 {
			outcomeSum2 += rpsOutcomes2[line]
		} else {
			fmt.Println("ALL DONE!")
		}
	}

	fmt.Println(outcomeSum2)
}
