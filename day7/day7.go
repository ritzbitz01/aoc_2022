package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findDirectories(part int) int {
	readFile, err := os.Open("day7_input_ex.txt")

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

	// Object = ["dirName"]["parent", "dir/filenames..."]interface{}

	type Directory struct {
		name string
		parent *Directory
		fileSizes []int
		subDirs map[string]*Directory
	}

	// filesystem := make(map[string]Directory)
	var filesystem Directory
	var currentDirectory Directory

	for _, line := range fileLines {
		if strings.HasPrefix(line,"$") {
			fmt.Println("COMMAND: ", line)
			if line[2:4] == "cd" {
				cdDirectory := line[5:]
				if cdDirectory == "/" {
					dir := Directory{name: cdDirectory, parent: nil, subDirs: make(map[string]*Directory), fileSizes: []int{}}
					filesystem = dir
					currentDirectory = dir
				} else if cdDirectory == ".." {
					fmt.Println(".. CUR DIR: ", currentDirectory)
					currentDirectory = *currentDirectory.parent
					fmt.Println(".. PAR DIR: ", currentDirectory)
				} else {
					fmt.Println("CURRENT DIRECTORY: ", currentDirectory)
					currentDirectory = *currentDirectory.subDirs[cdDirectory]
				}
			} else {
				// ls command
				continue
			}
		} else if strings.HasPrefix(line, "dir") {
			fmt.Println("CURRENT DIRECTORY: ", currentDirectory)
			fmt.Println("DIRECTORY: ", line)
			newDir := Directory{name: line[4:], parent: &currentDirectory, subDirs: make(map[string]*Directory), fileSizes: []int{}}
			currentDirectory.subDirs[line[4:]] = &newDir
		} else {
			fmt.Println("FILE: ", line)
			fileDetails := strings.Split(line, " ")
			fileSize, _ := strconv.Atoi(fileDetails[0])
			currentDirectory.fileSizes = append(currentDirectory.fileSizes, fileSize)
		}
	}
	fmt.Println(filesystem)
	return -1
}

func main() {
	fmt.Println("PART 1 INDEX: ", findDirectories(1))
	// fmt.Println("PART 2 INDEX: ", findDirectories(2))
}
