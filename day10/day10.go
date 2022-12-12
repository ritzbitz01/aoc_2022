package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMarker(part int) int {
	readFile, err := os.Open("day10_input.txt")
	// readFile, err := os.Open("day10_input_ex.txt")

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

	crtRow := 0
	crt := make(map[int][40]string)
	var crtPixel [40]string
	crtPosition := 0

	for _, line := range fileLines {
		input_vals := strings.Split(line, " ")
		registerArray := []int{xRegister-1, xRegister, xRegister+1}
		if input_vals[0] == "noop" {
			cycleNum ++
			if cycleNum == 20 || cycleNum == 60 || cycleNum == 100 || cycleNum == 140 || cycleNum == 180 || cycleNum == 220 {
				fmt.Println("NOOP ", cycleNum, "th Cycle Value: ", xRegister*cycleNum)
			}
			if contains(registerArray, crtPosition) {
				crtPixel[crtPosition] = "#"
			} else {
				crtPixel[crtPosition] = "."
			}
			if crtPosition == 39 {
				crt[crtRow] = crtPixel
				crtPosition = 0
				crtRow++
			} else {
				crtPosition++
			}
			
		} else {
			// if cycleNum == 20 || cycleNum == 60 || cycleNum == 100 || cycleNum == 140 || cycleNum == 180 || cycleNum == 220 {
			// 	fmt.Println("ADDX 1: ", cycleNum, "th Cycle Value: ", xRegister*cycleNum)
			// }
			cycleNum ++ 
			if cycleNum == 20 || cycleNum == 60 || cycleNum == 100 || cycleNum == 140 || cycleNum == 180 || cycleNum == 220  {
				fmt.Println("ADDX 2: ", cycleNum, "th Cycle Value: ", xRegister*cycleNum)
			}
			if contains(registerArray, crtPosition) {
				crtPixel[crtPosition] = "#"
			} else {
				crtPixel[crtPosition] = "."
			}
			if crtPosition == 39 {
				crt[crtRow] = crtPixel
				crtPosition = 0
				crtRow++
			} else {
				crtPosition++
			}
			cycleNum ++
			if cycleNum == 20 || cycleNum == 60 || cycleNum == 100 || cycleNum == 140 || cycleNum == 180 || cycleNum == 220  {
				fmt.Println("ADDX 3: ", cycleNum, "th Cycle Value: ", xRegister*cycleNum)
			}
			if contains(registerArray, crtPosition) {
				crtPixel[crtPosition] = "#"
			} else {
				crtPixel[crtPosition] = "."
			}
			if crtPosition == 39 {
				crt[crtRow] = crtPixel
				crtPosition = 0
				crtRow++
			} else {
				crtPosition++
			}
			regIncrease, _ := strconv.Atoi(input_vals[1])
			
			xRegister += regIncrease
		}
	}
	for _, c := range crt {
		fmt.Println(c)
	}
	fmt.Println(crt[0])
	fmt.Println(crt[1])
	fmt.Println(crt[2])
	fmt.Println(crt[3])
	fmt.Println(crt[4])
	fmt.Println(crt[5])
	return -1
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func main() {
	fmt.Println("PART 1 INDEX: ", findMarker(1))
	// fmt.Println("PART 2 INDEX: ", findMarker(2))
}
