package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var to_number = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func number(line string) int {
	var first string
	var last string
	for _, char := range line {
		if unicode.IsDigit(char) {
			if first == "" {
				first = string(char)
			} else {
				last = string(char)
			}
		}
	}
	if last == "" {
		last = first
	}
	n, _ := strconv.Atoi(first + last)
	return n
}

func main() {
	file, _ := os.Open("day-1/example.txt")
	scanner := bufio.NewScanner(file)
	sum_one := 0
	sum_two := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Part one
		sum_one += number(line)

		// Part two
		for word, digit := range to_number {
			// "two" -> "t2o" so "xtwone3four" becomes "xt2o1e3f4r" not "xtw134" (missing the "two")
			line = strings.ReplaceAll(line, word, word[:1]+digit+word[len(word)-1:])
		}
		sum_two += number(line)
	}

	fmt.Println("Part one:", sum_one)
	fmt.Println("Part two:", sum_two)
}
