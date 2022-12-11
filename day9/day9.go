package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func findMarker(part int) int {
	readFile, err := os.Open("day9_input.txt")
	// readFile, err := os.Open("day9_input_ex2.txt")

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

	tailPosVisitedTotal := 0

	// Track the positions tail has been
	tailPositions := make(map[string]bool)

	allTailPositions := []Position{{x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}}

	tailPositions["0_0"] = true
	tailPosVisitedTotal++
	for _, line := range fileLines {
		vals := strings.Split(line, " ")
		direction := vals[0]
		numSteps, _ := strconv.Atoi(vals[1])
		switch direction {
		case "U":
			for i := 0; i < numSteps; i++ {
				allTailPositions[0].y++
				for j := 1; j < len(allTailPositions); j++ {
					tailMoved := false
					if math.Abs(float64(allTailPositions[j-1].y-allTailPositions[j].y)) == 2 &&
						math.Abs(float64(allTailPositions[j-1].x-allTailPositions[j].x)) == 2 {
						allTailPositions[j].x = (allTailPositions[j-1].x + allTailPositions[j].x) / 2
						allTailPositions[j].y = (allTailPositions[j-1].y + allTailPositions[j].y) / 2
						tailMoved = true
					}
					if math.Abs(float64(allTailPositions[j-1].y-allTailPositions[j].y)) == 2 {
						if allTailPositions[j].y < allTailPositions[j-1].y {
							allTailPositions[j].y++
						} else {
							allTailPositions[j].y--
						}
						if allTailPositions[j-1].x != allTailPositions[j].x {
							allTailPositions[j].x = allTailPositions[j-1].x
						}
						tailMoved = true
					}
					if math.Abs(float64(allTailPositions[j-1].x-allTailPositions[j].x)) == 2 {
						if allTailPositions[j].x < allTailPositions[j-1].x {
							allTailPositions[j].x++
						} else {
							allTailPositions[j].x--
						}
						if allTailPositions[j-1].y != allTailPositions[j].y {
							allTailPositions[j].y = allTailPositions[j-1].y
						}
						tailMoved = true
					}
					if tailMoved {
						if j == 9 {
							if !tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] {
								tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] = true
								tailPosVisitedTotal++
							}
						}
					}
				}
			}
		case "R":
			for i := 0; i < numSteps; i++ {
				allTailPositions[0].x++
				for j := 1; j < len(allTailPositions); j++ {
					tailMoved := false
					if math.Abs(float64(allTailPositions[j-1].y-allTailPositions[j].y)) == 2 &&
						math.Abs(float64(allTailPositions[j-1].x-allTailPositions[j].x)) == 2 {
						allTailPositions[j].x = (allTailPositions[j-1].x + allTailPositions[j].x) / 2
						allTailPositions[j].y = (allTailPositions[j-1].y + allTailPositions[j].y) / 2
						tailMoved = true
					}
					if math.Abs(float64(allTailPositions[j-1].y-allTailPositions[j].y)) == 2 {
						if allTailPositions[j].y < allTailPositions[j-1].y {
							allTailPositions[j].y++
						} else {
							allTailPositions[j].y--
						}
						if allTailPositions[j-1].x != allTailPositions[j].x {
							allTailPositions[j].x = allTailPositions[j-1].x
						}
						tailMoved = true
					}
					if math.Abs(float64(allTailPositions[j-1].x-allTailPositions[j].x)) == 2 {
						if allTailPositions[j].x < allTailPositions[j-1].x {
							allTailPositions[j].x++
						} else {
							allTailPositions[j].x--
						}
						if allTailPositions[j-1].y != allTailPositions[j].y {
							allTailPositions[j].y = allTailPositions[j-1].y
						}
						tailMoved = true
					}
					if tailMoved {
						if j == 9 {
							if !tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] {
								tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] = true
								tailPosVisitedTotal++
							}
						}
					}
				}
			}
		case "D":
			for i := 0; i < numSteps; i++ {
				allTailPositions[0].y--
				for j := 1; j < len(allTailPositions); j++ {
					tailMoved := false
					if math.Abs(float64(allTailPositions[j-1].y-allTailPositions[j].y)) == 2 &&
						math.Abs(float64(allTailPositions[j-1].x-allTailPositions[j].x)) == 2 {
						allTailPositions[j].x = (allTailPositions[j-1].x + allTailPositions[j].x) / 2
						allTailPositions[j].y = (allTailPositions[j-1].y + allTailPositions[j].y) / 2
						tailMoved = true
					}
					if math.Abs(float64(allTailPositions[j-1].y-allTailPositions[j].y)) == 2 {
						if allTailPositions[j].y < allTailPositions[j-1].y {
							allTailPositions[j].y++
						} else {
							allTailPositions[j].y--
						}
						if allTailPositions[j-1].x != allTailPositions[j].x {
							allTailPositions[j].x = allTailPositions[j-1].x
						}
						tailMoved = true
					}
					if math.Abs(float64(allTailPositions[j-1].x-allTailPositions[j].x)) == 2 {
						if allTailPositions[j].x < allTailPositions[j-1].x {
							allTailPositions[j].x++
						} else {
							allTailPositions[j].x--
						}
						if allTailPositions[j-1].y != allTailPositions[j].y {
							allTailPositions[j].y = allTailPositions[j-1].y
						}
						tailMoved = true
					}
					if tailMoved {
						if j == 9 {
							if !tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] {
								tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] = true
								tailPosVisitedTotal++
							}
						}
					}
				}
			}
		case "L":
			for i := 0; i < numSteps; i++ {
				allTailPositions[0].x--
				for j := 1; j < len(allTailPositions); j++ {
					tailMoved := false
					if math.Abs(float64(allTailPositions[j-1].y-allTailPositions[j].y)) == 2 &&
						math.Abs(float64(allTailPositions[j-1].x-allTailPositions[j].x)) == 2 {
						allTailPositions[j].x = (allTailPositions[j-1].x + allTailPositions[j].x) / 2
						allTailPositions[j].y = (allTailPositions[j-1].y + allTailPositions[j].y) / 2
						tailMoved = true
					}
					if math.Abs(float64(allTailPositions[j-1].x-allTailPositions[j].x)) == 2 {
						if allTailPositions[j].x < allTailPositions[j-1].x {
							allTailPositions[j].x++
						} else {
							allTailPositions[j].x--
						}
						if allTailPositions[j-1].y != allTailPositions[j].y {
							allTailPositions[j].y = allTailPositions[j-1].y
						}
						tailMoved = true
					}
					if math.Abs(float64(allTailPositions[j-1].y-allTailPositions[j].y)) == 2 {
						if allTailPositions[j].y < allTailPositions[j-1].y {
							allTailPositions[j].y++
						} else {
							allTailPositions[j].y--
						}
						if allTailPositions[j-1].x != allTailPositions[j].x {
							allTailPositions[j].x = allTailPositions[j-1].x
						}
						tailMoved = true
					}
					if tailMoved {
						if j == 9 {
							if !tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] {
								tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] = true
								tailPosVisitedTotal++
							}
						}
					}
				}
			}
		default:
			fmt.Println("UNKNOWN DIRECTION: ", direction)
		}
	}

	return tailPosVisitedTotal
}

func main() {
	fmt.Println("PART 1 INDEX: ", findMarker(1))
	// fmt.Println("PART 2 INDEX: ", findMarker(2))
}
