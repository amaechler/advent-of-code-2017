package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("day9_input")
	if err != nil {
		log.Fatal(err)
	}

	totalScore, totalGarbageCount := CalculateStreamScores(content)

	fmt.Println("Total Group Score:", totalScore)
	fmt.Println("Total Garbage Count:", totalGarbageCount)
}

// CalculateStreamScores ...
func CalculateStreamScores(stream []byte) (int, int) {
	totalGroupScore := 0
	totalGarbageCount := 0
	currentGroupCount := 0
	isGarbage := false

	for i := 0; i < len(stream); i++ {
		switch string(stream[i]) {
		case "!":
			i++
			continue
		case "<":
			if isGarbage {
				totalGarbageCount++
			} else {
				isGarbage = true
			}
		case ">":
			isGarbage = false
		case "{":
			if isGarbage {
				totalGarbageCount++
			} else {
				currentGroupCount++
				totalGroupScore += currentGroupCount
			}
		case "}":
			if isGarbage {
				totalGarbageCount++
			} else {
				currentGroupCount--
			}
		default:
			if isGarbage {
				totalGarbageCount++
			}
		}
	}

	return totalGroupScore, totalGarbageCount
}
