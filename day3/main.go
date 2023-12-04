package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

var direction = [8][2]int{
	{-1, 0},
	{1, 0},
	{0, 1},
	{1, -1},
	{0, -1},
	{-1, -1},
	{-1, 1},
	{1, 1},
}

func main() {

	partOne()
}

func partOne() {
	file, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(file), "\n")

	symbolLocation := make(map[pos]bool)

	x := 1

	for _, line := range lines {

		for y, c := range line {

			if c == 46 {
				continue
			}

			if c >= 48 && c <= 57 {
				continue
			}

			position := pos{x, y}

			symbolLocation[position] = true
			// fmt.Println(position)

		}

		x++
		// break
	}

	x = 1

	totalSum := 0

	for _, line := range lines {

		isNumberValid := false
		currentNumber := "0"
		// fmt.Println(line)
		for y, c := range line {

			if c >= 48 && c <= 57 {

				currentNumber += string(c)

				if !isNumberValid {
					isNumberValid = isValid(pos{x, y}, symbolLocation)
				}
				// fmt.Println(string(c), isNumberValid)

			} else {

				if isNumberValid && currentNumber != "0" {
					sum, _ := strconv.Atoi(currentNumber)
					// fmt.Println(x, sum)
					totalSum += sum
				}
				isNumberValid = false
				currentNumber = "0"
			}
		}
		if isNumberValid && currentNumber != "0" {
			sum, _ := strconv.Atoi(currentNumber)
			// fmt.Println(x, sum)
			totalSum += sum
		}
		x++
		// if x > 1 {
		// 	break
		// }
	}

	fmt.Println(totalSum)

}

func isValid(position pos, symbolLocation map[pos]bool) bool {

	for _, dir := range direction {
		x, y := position.x+dir[0], position.y+dir[1]

		if symbolLocation[pos{x, y}] {
			// fmt.Println("checked true")
			return true
		}

	}

	return false
}
