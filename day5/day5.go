package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("day5_input.txt")
	// readFile, err := os.Open("day5_input_ex.txt")

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

	// var stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9 []string
	stacks := make(map[int][]string)
	lineNum := 0
	for _, line := range fileLines {
		if lineNum < 8 {
			i := 1
			idx := 1
			for i <= 9 {
				if strings.TrimSpace(string(line[idx])) != "" {
					// fmt.Println("I: ", i, " IDX: ", idx, " VAL: ", string(line[idx]))
					stacks[i] = append(stacks[i], string(line[idx]))
				}
				i++
				idx += 4
			}
		}

		// if lineNum < 3 {
		// 	i := 1
		// 	idx := 1
		// 	for i <= 3 {
		// 		if strings.TrimSpace(string(line[idx])) != "" {
		// 			// fmt.Println("I: ", i, " IDX: ", idx, " VAL: ", string(line[idx]))
		// 			stacks[i] = append(stacks[i], string(line[idx]))
		// 		}
		// 		i++
		// 		idx += 4
		// 	}
		// }

		if lineNum > 11 {
			// move 1 from 8 to 1
			fmt.Println("LINENUM: ", lineNum, "LINE: ", line)
			vals := strings.Split(line, " ")
			quantity, _ := strconv.Atoi(vals[1])
			source, _ := strconv.Atoi(vals[3])
			dest, _ := strconv.Atoi(vals[5])
			// fmt.Println("QUANTITY: ", quantity, "SOURCE: ", source, "DEST: ", dest)
			if lineNum == 54 {
				fmt.Println(stacks[1])
				fmt.Println(stacks[3])
			}
			i:=0
			for i < quantity {
				crate := ""
				crate, stacks[source] = stacks[source][0], stacks[source][1:]
				stacks[dest] = append(stacks[dest], crate)
				i++
			}
		}

		// if lineNum > 4 {
		// 	// move 1 from 8 to 1
		// 	fmt.Println("LINE: ", line)
		// 	vals := strings.Split(line, " ")
		// 	quantity, _ := strconv.Atoi(vals[1])
		// 	source, _ := strconv.Atoi(vals[3])
		// 	dest, _ := strconv.Atoi(vals[5])
		// 	// fmt.Println("QUANTITY: ", quantity, "SOURCE: ", source, "DEST: ", dest)
		// 	i:=0
		// 	for i < quantity {
		// 		crate := ""
		// 		crate, stacks[source] = stacks[source][0], stacks[source][1:]
		// 		stacks[dest] = append([]string{crate}, stacks[dest]...)
		// 		i++
		// 	}
		// 	fmt.Println("STACKS: ", stacks)
		// }
		lineNum++
	}

	fmt.Println("stack1: ", stacks);
	// fmt.Println("OVERLAP PAIRS: ", overlapPairs);

}
