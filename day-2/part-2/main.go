package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input")

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(666)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputArray []string
	var correctBoxIds []string

	for scanner.Scan() {
		inputArray = append(inputArray, scanner.Text())
	}

	for i := 0; i < len(inputArray)-1; i++ {
		for k := 0; k < len(inputArray); k++ {
			hasOneLetterDiff := false
			hasMoreThanOneLetterDiff := false
			for j := 0; j < len([]rune(inputArray[i])); j++ {
				if []rune(inputArray[i])[j] != []rune(inputArray[k])[j] {

					if hasOneLetterDiff {
						hasMoreThanOneLetterDiff = true
					} else {
						hasOneLetterDiff = true
					}
				}
			}

			if hasOneLetterDiff && !hasMoreThanOneLetterDiff {
				correctBoxIds = append(correctBoxIds, inputArray[i])
			}
		}
	}

	fmt.Println(correctBoxIds)
}
