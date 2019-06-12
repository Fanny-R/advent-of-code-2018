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

	two := 0
	three := 0

	for scanner.Scan() {
		occurences := countRuneOccurences(scanner.Text())

		hasTwo := false
		hasThree := false

		for _, value := range occurences {
			if value == 2 {
				hasTwo = true
			}

			if value == 3 {
				hasThree = true
			}
		}

		if hasTwo {
			two++
		}

		if hasThree {
			three++
		}
	}
	fmt.Println("Result: ", two*three)
}

func countRuneOccurences(line string) map[rune]int {
	result := map[rune]int{}

	for _, char := range line {
		result[char]++
	}

	return result
}
