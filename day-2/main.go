package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day-2/input.txt")
	scanner := bufio.NewScanner(file)
	game := 0
	total_one := 0
	total_two := 0

	for scanner.Scan() {
		line := scanner.Text()
		rounds := strings.Split(strings.Split(line, ": ")[1], ";")
		game++
		valid_rounds := 0
		max_red, max_green, max_blue := 0.0, 0.0, 0.0

		for _, round := range rounds {
			red := 12
			green := 13
			blue := 14
			cubes := strings.Split(round, ",")

			for _, cube := range cubes {
				tuple := strings.Split(strings.Trim(cube, " "), " ")
				n, _ := strconv.Atoi(tuple[0])
				switch tuple[1] {
				case "red":
					red -= n
					max_red = math.Max(max_red, float64(n))
				case "green":
					green -= n
					max_green = math.Max(max_green, float64(n))
				case "blue":
					blue -= n
					max_blue = math.Max(max_blue, float64(n))
				}
			}
			if red >= 0 && green >= 0 && blue >= 0 {
				valid_rounds++
			}
		}
		if valid_rounds == len(rounds) {
			total_one += game
		}
		total_two += int(max_red * max_green * max_blue)
	}

	fmt.Println("Part one:", total_one)
	fmt.Println("Part two:", total_two)
}
