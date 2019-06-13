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

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println("Line : ", line)
	}
}
