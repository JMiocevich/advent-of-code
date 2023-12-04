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

	// partOne("test.txt")
	partTwo("test.txt")
}

func partTwo(input string) {

	file, _ := os.ReadFile(input)

	lines := strings.Split(string(file), "\n")

	symbolLocation := make(map[pos]bool)

	x := 1

	for _, line := range lines {

		for y, c := range line {

			if c == 42 {
				position := pos{x, y}

				symbolLocation[position] = true

			}
		}
		x++

	}

	x = 1

	numberMap := make(map[pos]int)

	for _, line := range lines {
		currentNumber := "0"
		currentNumberLength := 0

		for y, c := range line {
			if c >= 48 && c <= 57 {
				currentNumber += string(c)
				currentNumberLength++
			} else {
				sum, _ := strconv.Atoi(currentNumber)

				currentNumberLength++
				if sum != 0 {
					for i := 0; i < currentNumberLength-1; i++ {

						position := pos{x, y - i}

						numberMap[position] = sum
					}
				}
				currentNumberLength = 0
				currentNumber = "0"

			}
		}

		if currentNumber != "0" {
			sum, _ := strconv.Atoi(currentNumber)
			if sum != 0 {
				// println(sum)
				position := pos{x, len(line)}
				numberMap[position] = sum

			}
		}
		x++
	}

	// for pos, val := range numberMap {
	// 	fmt.Println(pos, val)
	// }

	sum := 0

	for symbolPostion := range symbolLocation {

		valOne := 1
		valTwo := 1
		valThree := 1

		// check above

		for i := -1; i < 2; i++ {
			position := pos{symbolPostion.x - i, symbolPostion.y - 1}
			value := numberMap[position]

			if value != 0 {
				valOne = value
			}
		}

		for i := -1; i < 2; i++ {
			position := pos{symbolPostion.x - i, symbolPostion.y}
			value := numberMap[position]

			if value != 0 {
				valTwo = value
			}
		}

		for i := -1; i < 2; i++ {
			position := pos{symbolPostion.x - i, symbolPostion.y + 1}
			value := numberMap[position]

			if value != 0 {
				valThree = value
			}
		}

		fmt.Println(valOne, valTwo, valThree)
		localSum := valOne * valTwo * valThree

		sum += localSum
	}

	fmt.Println(sum)
}

func partOne(input string) {
	file, _ := os.ReadFile(input)

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
