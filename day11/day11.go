package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	itemList []big.Int
	operationOperator string
	operationVal int64
	testDivisibleBy int64
	ifTrue int
	ifFalse int
	inspectCount int
}

func findMarker(part int) int {
	// readFile, err := os.Open("day11_input.txt")
	readFile, err := os.Open("day11_input_ex.txt")

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

	monkeys := make(map[int]Monkey)
	lineNum := 0
	monkeyNum := 0

	itemListMonkey := []big.Int{}
	opOp := ""
	var opVal *big.Int
	var testDivBy big.Int
	trueThrowToMonkey := 0
	falseThrowToMonkey := 0
	for _, line := range fileLines {
		// fmt.Println("LINE ", lineNum, ": ", line)
		switch lineNum{
		case 0:
			lineNum++
			continue
		case 1:
			lineVals := strings.Split(line, ":")
			itemList := strings.Split(lineVals[1], ",")
			for _, n := range itemList {
				itemInt, err := strconv.Atoi(strings.TrimSpace(n))
				blah := int64(itemInt)
				itemN := big.NewInt(blah)
				if err != nil {
					fmt.Println("ERROR CONVERSION: ", err)
				}
				itemListMonkey = append(itemListMonkey, *itemN)
			}
			lineNum++
		case 2:
			vals := strings.Split(line, "=")
			valsVals := strings.Split(vals[1], " ")
			opOp = valsVals[2]
			opValInt, _ := strconv.Atoi(strings.TrimSpace(valsVals[3]))
			opVal = big.NewInt(opValInt)
			lineNum++
		case 3:
			vals := strings.Split(line, " ")
			// fmt.Println("DIV: ", vals[5])
			testDivByInt, _ := strconv.Atoi(strings.TrimSpace(vals[5]))
			testDivBy = int64(testDivByInt)
			lineNum++
		case 4:
			vals := strings.Split(line, " ")
			// fmt.Println("TRUE: ", vals[9])
			trueThrowToMonkey, _ = strconv.Atoi(strings.TrimSpace(vals[9]))
			lineNum++
		case 5:
			vals := strings.Split(line, " ")
			// fmt.Println("FALSE: ", vals[9])
			falseThrowToMonkey, _ = strconv.Atoi(strings.TrimSpace(vals[9]))
			lineNum++
			monkeys[monkeyNum] = Monkey{itemList: itemListMonkey, 
				operationOperator: opOp,
				operationVal: opVal,
				testDivisibleBy: testDivBy,
				ifTrue: trueThrowToMonkey,
				ifFalse: falseThrowToMonkey,
				inspectCount: 0}

			itemListMonkey = []big.Int{}
			opOp = ""
			opVal = 0
			testDivBy = 0
			trueThrowToMonkey = 0
			falseThrowToMonkey = 0
			monkeyNum++
			continue
		case 6:
			lineNum=0
			continue
		default:
			fmt.Println("WHY DID THIS HAPPEN??? ", lineNum)
		} 
	}

	for roundNum:=0; roundNum < 10000; roundNum++ {
		for i:=0; i < len(monkeys); i++ {
			// fmt.Println("Monkey ", i, ": ", monkeys[i]) 
			itemListItr := monkeys[i].itemList
			// fmt.Println("ITEM LIST ITR: ", itemListItr)
			for j:=0; j < len(itemListItr); j++ {
				// fmt.Println("ITEM LIST ITR INNER: ", itemListItr)
				itm := itemListItr[j]
				// fmt.Println("ITEM VAL: ", itm)
				var worryLevel big.Int
				var operVal big.Int
				if monkeys[i].operationVal == 0 {
					operVal = itm
				} else {
					operVal = monkeys[i].operationVal
				}			
				switch monkeys[i].operationOperator {
				case "+":
					worryLevel = itm + operVal
				case "*":
					worryLevel = itm * operVal
				default:
					fmt.Println("BAD OPERATOR")
				}
				// fmt.Println("WORRY LEVEL OP: ", worryLevel)
				// worryLevel = int(worryLevel/3)
				// fmt.Println("WORRY LEVEL INS: ", worryLevel)
				if worryLevel%monkeys[i].testDivisibleBy == 0 {
					// fmt.Println("THROWING ITEM TO: ", monkeys[i].ifTrue, " WLVL: ", worryLevel)
					if entry, ok := monkeys[monkeys[i].ifTrue]; ok {
						entry.itemList = append(monkeys[monkeys[i].ifTrue].itemList, worryLevel)
						monkeys[monkeys[i].ifTrue] = entry
					}
					if entry2, ok := monkeys[i]; ok {
						_, entry2.itemList = entry2.itemList[0], entry2.itemList[1:]
						// entry2.itemList = append(monkeys[i].itemList[:0], monkeys[i].itemList[1:]...)
						monkeys[i] = entry2
					}
				} else {
					// fmt.Println("THROWING ITEM TO: ", monkeys[i].ifFalse, " WLVL: ", worryLevel)
					if entry, ok := monkeys[monkeys[i].ifFalse]; ok {
						entry.itemList = append(monkeys[monkeys[i].ifFalse].itemList, worryLevel)
						monkeys[monkeys[i].ifFalse] = entry
					}
					if entry2, ok := monkeys[i]; ok {
						_, entry2.itemList = entry2.itemList[0], entry2.itemList[1:]
						// entry2.itemList = append(monkeys[i].itemList[:0], monkeys[i].itemList[1:]...)
						monkeys[i] = entry2
					}
					// fmt.Println("ITEM LIST ITR FINAL: ", itemListItr)
				}
				if entry, ok := monkeys[i]; ok {
					entry.inspectCount++
					monkeys[i] = entry
				}
			}
		}
		if roundNum%1000 == 0 {
			fmt.Println("----------ROUND ", roundNum, "---------------------")
			fmt.Println("MONKEY 0: ", monkeys[0])
			fmt.Println("MONKEY 1: ", monkeys[1])
			fmt.Println("MONKEY 2: ", monkeys[2])
			fmt.Println("MONKEY 3: ", monkeys[3])
			fmt.Println("MONKEY 4: ", monkeys[4])
			fmt.Println("MONKEY 5: ", monkeys[5])
			fmt.Println("MONKEY 6: ", monkeys[6])
			fmt.Println("MONKEY 7: ", monkeys[7])
		}
	}
	fmt.Println("MONKEY 0: ", monkeys[0])
	fmt.Println("MONKEY 1: ", monkeys[1])
	fmt.Println("MONKEY 2: ", monkeys[2])
	fmt.Println("MONKEY 3: ", monkeys[3])
	fmt.Println("MONKEY 4: ", monkeys[4])
	fmt.Println("MONKEY 5: ", monkeys[5])
	fmt.Println("MONKEY 6: ", monkeys[6])
	fmt.Println("MONKEY 7: ", monkeys[7])
	return -1
}

func main() {
	fmt.Println("PART 1 INDEX: ", findMarker(1))
	// fmt.Println("PART 2 INDEX: ", findMarker(2))
}
