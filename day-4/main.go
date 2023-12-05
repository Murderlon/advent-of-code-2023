package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	file, _ := os.Open("day-4/input.txt")
	scanner := bufio.NewScanner(file)
	total_one, total_two, game := 0, 0, 0
	cards := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		lists := strings.Split(strings.Split(line, ":")[1], "|")
		winning := strings.Fields(lists[0])
		numbers := strings.Fields(lists[1])
		points := 0.0
		wins := 0

		cards[game+1]++
		for _, win := range winning {
			if slices.Contains(numbers, win) {
				wins++
				points = math.Max(1.0, points*2)
				cards[game+1+wins] = cards[game+1+wins] + cards[game+1]
			}
		}

		total_one += int(points)
		game++
	}

	for _, v := range cards {
		total_two += v
	}

	fmt.Println("Part one:", total_one)
	fmt.Println("Part two:", total_two)
}
