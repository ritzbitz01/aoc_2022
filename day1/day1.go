package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	readFile, err := os.Open("day1_input.txt")

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

	var elfCalories []int
	var elfNum = 0
	var mealNum = 0
	var numElves = 0
	for _, line := range fileLines {
		// fmt.Println(line)
		if len(line) > 0 {
			intVar, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("ERROR WITH ELF ", elfNum, line, " is not a number.")
			}
			if mealNum == 0 {
				elfCalories = append(elfCalories, intVar)
			} else {
				elfCalories[elfNum] += intVar
			}
			mealNum++
		} else {
			fmt.Println("NEW ELF!")
			numElves++
			elfNum++
			mealNum = 0
		}
	}

	sort.Slice(elfCalories, func(i, j int) bool {
		return elfCalories[i] < elfCalories[j]
	})

	for _, elf := range elfCalories {
		fmt.Println(elf)
	}

	fmt.Println(elfCalories[len(elfCalories)-1])
	fmt.Println(elfCalories[len(elfCalories)-2])
	fmt.Println(elfCalories[len(elfCalories)-3])
	var topThreeElves = 0
	topThreeElves = elfCalories[len(elfCalories)-1] + elfCalories[len(elfCalories)-2] + elfCalories[len(elfCalories)-3]
	fmt.Println(topThreeElves)
}
