package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMarker(part int) int {
	readFile, err := os.Open("day6_input.txt")

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

	var charArray []string
	var numUnique int
	if part == 1 {
		numUnique = 4
	} else {
		numUnique = 14
	}

	for _, line := range fileLines {
		for idx, e := range line {
			if idx <= numUnique -1 {
				charArray = append(charArray, string(e))
			} else {
				if !findDuplicates(charArray) {
					return idx
				} else {
					_, charArray = charArray[0], charArray[1:]
					charArray = append(charArray, string(e))
				}
			}

		}
		
	}
	return -1
}

func findDuplicates(chars []string) bool{
	visited := make(map[string]bool, 0)
	for i:=0; i<len(chars); i++{
	   if visited[chars[i]] == true{
		  return true
	   } else {
		  visited[chars[i]] = true
	   }
	}
	return false
}

func main() {
	fmt.Println("PART 1 INDEX: ", findMarker(1))
	fmt.Println("PART 2 INDEX: ", findMarker(2))
}
