package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Tree struct {
	height  int
	visible bool
}

func findMarker(part int) int {
	readFile, err := os.Open("day9_input.txt")
	// readFile, err := os.Open("day9_input_ex.txt")

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

	numTreesVisible := 0
	maxScenicScore := 0

	// Create the tree array
	var treeArray [99][99]Tree
	// var treeArray [5][5]Tree
	rowNum := 0
	for _, line := range fileLines {
		for idx, e := range line {
			treeHeight, _ := strconv.Atoi(string(e))
			tree := Tree{height: treeHeight, visible: false}
			treeArray[rowNum][idx] = tree
		}
		rowNum++
	}

	if part == 1 {
		// Iterate the tree array
		for idx, row := range treeArray {
			if idx == 0 || idx == len(treeArray)-1 {
				for i := 0; i < len(row); i++ {
					treeArray[idx][i].visible = true
					numTreesVisible++
				}
				continue
			}

			numTreesVisible++
			treeArray[idx][0].visible = true
			for i := 1; i < len(row)-1; i++ {
				// row
				leftRow := append([]Tree{}, row[0:i]...)
				sort.Slice(leftRow, func(k, l int) bool {
					return leftRow[k].height < leftRow[l].height
				})

				if leftRow[len(leftRow)-1].height < row[i].height {
					treeArray[idx][i].visible = true
					numTreesVisible++
					continue
				}

				rightRow := append([]Tree{}, row[i+1:]...)
				sort.Slice(rightRow, func(k, l int) bool {
					return rightRow[k].height < rightRow[l].height
				})

				if rightRow[len(rightRow)-1].height < row[i].height {
					treeArray[idx][i].visible = true
					numTreesVisible++
					continue
				}

				// column
				var topColumn []Tree
				for ii := 0; ii < idx; ii++ {
					topColumn = append(topColumn, treeArray[ii][i])
				}
				sort.Slice(topColumn, func(k, l int) bool {
					return topColumn[k].height < topColumn[l].height
				})

				if topColumn[len(topColumn)-1].height < row[i].height {
					treeArray[idx][i].visible = true
					numTreesVisible++
					continue
				}

				var bottomColumn []Tree
				for ii := idx + 1; ii < len(treeArray); ii++ {
					bottomColumn = append(bottomColumn, treeArray[ii][i])
				}
				sort.Slice(bottomColumn, func(k, l int) bool {
					return bottomColumn[k].height < bottomColumn[l].height
				})

				if bottomColumn[len(bottomColumn)-1].height < row[i].height {
					treeArray[idx][i].visible = true
					numTreesVisible++
					continue
				}

			}
			numTreesVisible++
			treeArray[idx][len(row)-1].visible = true
		}
		return numTreesVisible
	} else {

		// Iterate the tree array
		for idx, row := range treeArray {
			if idx == 0 || idx == len(treeArray)-1 {
				continue
			}
			for i := 1; i < len(row)-1; i++ {
				// row
				leftTreesSeen := 0
				leftRow := append([]Tree{}, row[0:i]...)

				for j := i - 1; j >= 0; j-- {
					if leftRow[j].height < row[i].height {
						leftTreesSeen++
					} else {
						leftTreesSeen++
						break
					}
				}

				rightTreesSeen := 0
				rightRow := append([]Tree{}, row[i+1:]...)

				for j := 0; j <= len(rightRow)-1; j++ {
					if rightRow[j].height < row[i].height {
						rightTreesSeen++
					} else {
						rightTreesSeen++
						break
					}
				}

				// column
				topTreesSeen := 0
				var topColumn []Tree
				for ii := 0; ii < idx; ii++ {
					topColumn = append(topColumn, treeArray[ii][i])
				}

				for j := idx - 1; j >= 0; j-- {
					if topColumn[j].height < row[i].height {
						topTreesSeen++
					} else {
						topTreesSeen++
						break
					}
				}

				bottomTreesSeen := 0
				var bottomColumn []Tree
				for ii := idx + 1; ii < len(treeArray); ii++ {
					bottomColumn = append(bottomColumn, treeArray[ii][i])
				}
				for j := 0; j <= len(bottomColumn)-1; j++ {
					if bottomColumn[j].height < row[i].height {
						bottomTreesSeen++
					} else {
						bottomTreesSeen++
						break
					}
				}
				treeScenicScore := leftTreesSeen * rightTreesSeen * topTreesSeen * bottomTreesSeen
				if treeScenicScore > maxScenicScore {
					maxScenicScore = treeScenicScore
				}
			}

		}
		return maxScenicScore
	}

}

func main() {
	fmt.Println("PART 1 INDEX: ", findMarker(1))
	fmt.Println("PART 2 INDEX: ", findMarker(2))
}
