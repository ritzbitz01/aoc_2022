package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func organizeStacks(part int) string {
	readFile, err := os.Open("day5_input.txt")

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

	stacks := make(map[int][]string)
	lineNum := 0
	for _, line := range fileLines {
		if lineNum < 8 {
			i := 1
			idx := 1
			for i <= 9 {
				if strings.TrimSpace(string(line[idx])) != "" {
					stacks[i] = append(stacks[i], string(line[idx]))
				}
				i++
				idx += 4
			}
		}

		// PART 1
		if part == 1 {
			if lineNum > 9 {
				vals := strings.Split(line, " ")
				quantity, _ := strconv.Atoi(vals[1])
				source, _ := strconv.Atoi(vals[3])
				dest, _ := strconv.Atoi(vals[5])

				ii := 0
				for ii < quantity {
					crate := ""
					crate, stacks[source] = stacks[source][0], stacks[source][1:]
					stacks[dest] = append([]string{crate}, stacks[dest]...)
					ii++
				}
			}
		}

		if part == 2 {
			// PART 2
			if lineNum > 9 {
				vals := strings.Split(line, " ")
				quantity, _ := strconv.Atoi(vals[1])
				source, _ := strconv.Atoi(vals[3])
				dest, _ := strconv.Atoi(vals[5])

				var crates []string

				crates, stacks[source] = stacks[source][0:quantity], stacks[source][quantity:]

				for i := len(crates) - 1; i >= 0; i-- {
					stacks[dest] = append([]string{crates[i]}, stacks[dest]...)
				}
			}
		}
		lineNum++
	}

	topStacks := ""
	for i := 1; i < 10; i++ {
		topStacks += stacks[i][0]
	}
	return topStacks

}

func main() {
	fmt.Println(organizeStacks(1))
	fmt.Println(organizeStacks(2))
}
