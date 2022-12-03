package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("day3_input.txt")

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

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	prioritySum := 0

	for _, line := range fileLines {
		firstSet := line[0:(len(line)/2)]
		secondSet := line[len(line)/2:]

		hash := make(map[string]bool)
		for _, e := range firstSet {
			hash[string(e)] = true
		}
		for _, e := range secondSet {
			if hash[string(e)] {
				index := strings.Index(letters, string(e)) + 1
				prioritySum += index
				break
			}
		}
	}

	fmt.Println("PART 1 SUM: ", prioritySum)

	prioritySum = 0

	elfNum := 0
	elfOne := ""
	elfTwo := ""
	elfThree := ""
	for _, line := range fileLines {
		switch elfNum {
		case 0:
			elfOne = line
		case 1:
			elfTwo = line
		case 2:
			elfThree = line
		}

		if elfNum == 2 {
			var firstInter  []string

			hash := make(map[string]bool)
			for _, e := range elfOne {
				hash[string(e)] = true
			}
			for _, e := range elfTwo {
				if hash[string(e)] {
					firstInter = append(firstInter, string(e))
				}
			}

			hash2 := make(map[string]bool)
			for _, e := range firstInter {
				hash2[string(e)] = true
			}
			for _, e := range elfThree {
				if hash2[string(e)] {
					index := strings.Index(letters, string(e)) + 1
					prioritySum += index
					break
				}
			}
			elfNum = 0
		} else {
			elfNum ++
		}
	}

	fmt.Println("PART 2 SUM: ", prioritySum)

}
