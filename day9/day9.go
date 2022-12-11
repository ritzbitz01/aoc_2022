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
	// readFile, err := os.Open("day9_input.txt")
	readFile, err := os.Open("day9_input_ex2.txt")

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
	// var tailPositions []Position
	tailPositions := make(map[string]bool)

	// headPosition := Position{x: 0, y: 0}
	// tailPosition := Position{x: 0, y: 0}
	allTailPositions := []Position{{x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}}

	// fmt.Println("TAIL POS: ", tailPosition.x, " ", tailPosition.y)
	tailPositions["0_0"] = true
	tailPosVisitedTotal++
	for _, line := range fileLines {
		vals := strings.Split(line, " ")
		direction := vals[0]
		numSteps, _ := strconv.Atoi(vals[1])
		switch direction {
		case "U":
			fmt.Println("MOVE UP: ", numSteps)
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
						if i == 9 {
							if !tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] {
								fmt.Println("TAIL POS: ", allTailPositions[j].x, " ", allTailPositions[j].y)
								tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] = true
								tailPosVisitedTotal++
							}
						}
					}
				}
			}
		case "R":
			fmt.Println("MOVE RIGHT: ", numSteps)
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
						if i == 9 {
							if !tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] {
								fmt.Println("TAIL POS: ", allTailPositions[j].x, " ", allTailPositions[j].y)
								tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] = true
								tailPosVisitedTotal++
							}
						}
					}
				}
			}
		case "D":
			fmt.Println("MOVE DOWN: ", numSteps)
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
						if i == 9 {
							if !tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] {
								fmt.Println("TAIL POS: ", allTailPositions[j].x, " ", allTailPositions[j].y)
								tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] = true
								tailPosVisitedTotal++
							}
						}
					}
				}
			}
		case "L":
			fmt.Println("MOVE LEFT: ", numSteps)
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
								fmt.Println("TAIL POS: ", allTailPositions[j].x, " ", allTailPositions[j].y)
								tailPositions[fmt.Sprintf("%d_%d", allTailPositions[j].x, allTailPositions[j].y)] = true
								tailPosVisitedTotal++
							}
						}
					}
				}
				// fmt.Println("------- STEP ", i, " --------")
				// fmt.Println("HEAD POSITION: ", allTailPositions[0].x, " ", allTailPositions[0].y)
				// fmt.Println("SECOND POSITION: ", allTailPositions[1].x, " ", allTailPositions[1].y)
				// fmt.Println("THIRD POSITION: ", allTailPositions[2].x, " ", allTailPositions[2].y)
				// fmt.Println("FOURTH POSITION: ", allTailPositions[3].x, " ", allTailPositions[3].y)
				// fmt.Println("FIFTH POSITION: ", allTailPositions[4].x, " ", allTailPositions[4].y)
				// fmt.Println("SIXTH POSITION: ", allTailPositions[5].x, " ", allTailPositions[5].y)
				// fmt.Println("SEVENTH POSITION: ", allTailPositions[6].x, " ", allTailPositions[6].y)
				// fmt.Println("EIGHTH POSITION: ", allTailPositions[7].x, " ", allTailPositions[7].y)
				// fmt.Println("NINTH POSITION: ", allTailPositions[8].x, " ", allTailPositions[8].y)
				// fmt.Println("TAIL POSITION: ", allTailPositions[9].x, " ", allTailPositions[9].y)
			}
		default:
			fmt.Println("UNKNOWN DIRECTION: ", direction)
		}
		fmt.Println("HEAD POSITION: ", allTailPositions[0].x, " ", allTailPositions[0].y)
		fmt.Println("SECOND POSITION: ", allTailPositions[1].x, " ", allTailPositions[1].y)
		fmt.Println("THIRD POSITION: ", allTailPositions[2].x, " ", allTailPositions[2].y)
		fmt.Println("FOURTH POSITION: ", allTailPositions[3].x, " ", allTailPositions[3].y)
		fmt.Println("FIFTH POSITION: ", allTailPositions[4].x, " ", allTailPositions[4].y)
		fmt.Println("SIXTH POSITION: ", allTailPositions[5].x, " ", allTailPositions[5].y)
		fmt.Println("SEVENTH POSITION: ", allTailPositions[6].x, " ", allTailPositions[6].y)
		fmt.Println("EIGHTH POSITION: ", allTailPositions[7].x, " ", allTailPositions[7].y)
		fmt.Println("NINTH POSITION: ", allTailPositions[8].x, " ", allTailPositions[8].y)
		fmt.Println("TAIL POSITION: ", allTailPositions[9].x, " ", allTailPositions[9].y)
	}

	fmt.Println("TAIL POSITIONS: ", tailPositions)

	return tailPosVisitedTotal

}

func main() {
	fmt.Println("PART 1 INDEX: ", findMarker(1))
	// fmt.Println("PART 2 INDEX: ", findMarker(2))
}
