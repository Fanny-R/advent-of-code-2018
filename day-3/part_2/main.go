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
	id           int
	x            int
	y            int
	width        int
	height       int
	squareInches []squareInch
}

type squareInch []*claim

func main() {
	file, err := os.Open("./input2")

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(666)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fabric [1000][1000]squareInch

	var claims []*claim

	for scanner.Scan() {
		line := scanner.Text()

		claim := claimFromLine(line)
		for i := claim.x; i < claim.x+claim.width; i++ {
			for j := claim.y; j < claim.y+claim.height; j++ {
				fabric[i][j] = append(fabric[i][j], claim)

				claim.squareInches = append(claim.squareInches, fabric[i][j])
			}
		}
		claims = append(claims, claim)
	}

	for _, claim := range claims {
		var lengthInches []int
		for _, squareInch := range claim.squareInches {
			lengthInches = append(lengthInches, len(squareInch))
		}
		fmt.Println("Result : ", claim.id, lengthInches)

		if claim.isAlone() {
			fmt.Println("Result : ", claim.id)
			// return
		}
	}
}

func (c *claim) isAlone() bool {
	for _, squareInch := range c.squareInches {
		if len(squareInch) > 1 {
			return false
		}
	}
	return true
}

func claimFromLine(line string) *claim {
	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	result := strings.FieldsFunc(line, f)

	return &claim{
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
