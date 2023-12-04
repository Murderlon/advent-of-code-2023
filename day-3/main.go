package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

type Symbol struct {
	value   string
	matches []int
	x, y    int
}

type Number struct {
	value     string
	x1, x2, y int
}

func (s *Symbol) isCollinear(n *Number) bool {
	for i := n.x1; i <= n.x2; i++ {
		dx := float64(s.x - i)
		dy := float64(s.y - n.y)
		if math.Abs(dy) <= 1 && math.Abs(dx) <= 1 {
			return true
		}
	}

	return false
}

func (s *Symbol) isGear() bool {
	if s.value == "*" && len(s.matches) == 2 {
		return true
	}
	return false
}

func main() {
	file, _ := os.Open("day-3/input.txt")
	scanner := bufio.NewScanner(file)
	numbers := make([]Number, 0)
	symbols := make([]Symbol, 0)
	y, sum_one, sum_two := 0, 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		number := Number{x1: -1}

		for x, char := range line {
			if unicode.IsDigit(char) {
				number.value += string(char)
				number.y = y
				number.x2 = x
				if number.x1 == -1 {
					number.x1 = x
				}
				if x == len(line)-1 || !unicode.IsDigit(rune(line[x+1])) {
					numbers = append(numbers, number)
					number = Number{x1: -1}
				}
			} else if char != '.' {
				symbols = append(symbols, Symbol{value: string(char), x: x, y: y})
			}
		}
		y++
	}

	for _, symbol := range symbols {
		for _, number := range numbers {
			if symbol.isCollinear(&number) {
				n, _ := strconv.Atoi(number.value)
				symbol.matches = append(symbol.matches, n)
				sum_one += n
			}
		}

		if symbol.isGear() {
			sum_two += symbol.matches[0] * symbol.matches[1]
		}
	}

	fmt.Println("Part one:", sum_one)
	fmt.Println("Part two:", sum_two)
}
