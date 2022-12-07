package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Directory struct {
	name      string
	parent    *Directory
	fileSizes []int
	subDirs   map[string]*Directory
}

var dirsLess100k = 0
var dirsLess100kSize = 0
var allDirs []int

func findDirectories() int {
	readFile, err := os.Open("day7_input.txt")

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

	var filesystem *Directory
	var currentDirectory *Directory

	for _, line := range fileLines {
		if strings.HasPrefix(line, "$") {
			if line[2:4] == "cd" {
				cdDirectory := line[5:]
				if cdDirectory == "/" {
					dir := Directory{name: cdDirectory, parent: nil, subDirs: make(map[string]*Directory), fileSizes: []int{}}
					filesystem = &dir
					currentDirectory = &dir
				} else if cdDirectory == ".." {
					currentDirectory = currentDirectory.parent
				} else {
					currentDirectory = currentDirectory.subDirs[cdDirectory]
				}
			} else {
				// ls command
				continue
			}
		} else if strings.HasPrefix(line, "dir") {
			newDir := Directory{name: line[4:], parent: currentDirectory, subDirs: make(map[string]*Directory), fileSizes: []int{}}
			currentDirectory.subDirs[line[4:]] = &newDir
		} else {
			fileDetails := strings.Split(line, " ")
			fileSize, _ := strconv.Atoi(fileDetails[0])
			currentDirectory.fileSizes = append(currentDirectory.fileSizes, fileSize)
		}
	}

	totalSize := directorySize(*filesystem)

	return totalSize
}

func directorySize(dir Directory) int {
	dirSize := 0
	for _, subDir := range dir.subDirs {
		dirSize += directorySize(*subDir)
	}

	for _, fileSize := range dir.fileSizes {
		dirSize += fileSize
	}

	if dirSize <= 100000 {
		dirsLess100k++
		dirsLess100kSize += dirSize
	}

	allDirs = append(allDirs, dirSize)

	return dirSize
}

func main() {
	fileSystemSize := findDirectories()
	fmt.Println("100K Dir Size Total: ", dirsLess100kSize)

	remainingSpace := 70000000 - fileSystemSize
	spaceNeeded := 30000000 - remainingSpace

	sort.Ints(allDirs)
	for _, sz := range allDirs {
		if sz >= spaceNeeded {
			fmt.Println("Smallest size of deleteable dir: ", sz)
			break
		}
	}
}
