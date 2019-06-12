package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func main() {
	file, err := os.Open("./input")

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(666)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fabric [1000][1000]int

	for scanner.Scan() {
		line := scanner.Text()

		claim := claimFromLine(line)

		for i := claim.x; i < claim.x+claim.width; i++ {

			for j := claim.y; j < claim.y+claim.height; j++ {

				fabric[i][j]++
			}
		}
	}

	cptSquareInches := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] >= 2 {
				cptSquareInches++
			}
		}
	}

	fmt.Println("Result : ", cptSquareInches)
}

func claimFromLine(line string) claim {
	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	result := strings.FieldsFunc(line, f)

	return claim{
		id:     convStringToInt(result[0]),
		x:      convStringToInt(result[1]),
		y:      convStringToInt(result[2]),
		width:  convStringToInt(result[3]),
		height: convStringToInt(result[4]),
	}
}

func convStringToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(666)
	}

	return i
}
