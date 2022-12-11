package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMarker(part int) int {
	// readFile, err := os.Open("day10_input.txt")
	readFile, err := os.Open("day10_input_ex.txt")

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

	cycleNum := 0
	xRegister := 1

	for _, line := range fileLines {
		input_vals := strings.Split(line, " ")
		if input_vals[0] == "noop" {
			if cycleNum == 20 || cycleNum == 60 || cycleNum == 100 || cycleNum == 140 || cycleNum == 180 || cycleNum == 220 {
				fmt.Println(cycleNum, "th Cycle Value: ", xRegister*cycleNum)
			} else if cycleNum == 219 {
				fmt.Println(220, "th Cycle Value: ", xRegister*220)
			}
			cycleNum ++
		} else {
			if cycleNum == 20 || cycleNum == 60 || cycleNum == 100 || cycleNum == 140 || cycleNum == 180 || cycleNum == 220 {
				fmt.Println(cycleNum, "th Cycle Value: ", xRegister*cycleNum)
			}else if cycleNum == 219 {
				fmt.Println(220, "th Cycle Value: ", xRegister*220)
			}
			cycleNum ++ 
			if cycleNum == 20 || cycleNum == 60 || cycleNum == 100 || cycleNum == 140 || cycleNum == 180 || cycleNum == 220  {
				fmt.Println(cycleNum, "th Cycle Value: ", xRegister*cycleNum)
			}else if cycleNum == 219 {
				fmt.Println(220, "th Cycle Value: ", xRegister*220)
			}
			cycleNum ++
			regIncrease, _ := strconv.Atoi(input_vals[1])
			if cycleNum > 210 {
				fmt.Println("CYCLE: ", cycleNum, " X REGISTER INCREASE: ", regIncrease)
			}
			
			xRegister += regIncrease
		}
	}
	return -1
}

func main() {
	fmt.Println("PART 1 INDEX: ", findMarker(1))
	// fmt.Println("PART 2 INDEX: ", findMarker(2))
}
