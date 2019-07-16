package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
	"unicode"
)

type record struct {
	extractedTime time.Time
	minutes       string
	action        string
}

func main() {
	var records []*record

	file, err := os.Open("./example-input")

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(666)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		records = append(records, recordsFromLine(line))
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].extractedTime.Before(records[j].extractedTime)
	})

	for _, record := range records {
		fmt.Println("record : ", record.extractedTime)
	}
}

func recordsFromLine(line string) *record {
	f := func(c rune) bool {
		return !unicode.IsNumber(c) && !unicode.IsLetter(c)
	}

	result := strings.FieldsFunc(line, f)

	// use regex instead ?
	// re := regexp.MustCompile("\\[(.*?)\\]")
	// fmt.Println(re.FindAllString(line, -1))
	extractedTime, _ := time.Parse(time.RFC3339, result[0]+"-"+result[1]+"-"+result[2]+"T"+result[3]+":"+result[4]+":00Z")

	return &record{
		extractedTime: extractedTime,
		minutes:       result[4],
		// Guard # / asleep / up
		action: result[6],
	}
}
