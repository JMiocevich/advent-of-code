package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	res := partOne("input")

	fmt.Printf("Answer is part one : %d \n", res)

	res2 := partTwo("input")

	fmt.Printf("Answer is part one : %d \n", res2)
}

type race struct {
	time     int
	distance int
}

func partOne(input string) int {

	file, _ := os.ReadFile(input)
	lines := strings.Split(string(file), "\n")

	// Extract numbers from each line
	timeNumbers := extractNumbers(lines[0])
	distanceNumbers := extractNumbers(lines[1])

	// Check if timeNumbers and distanceNumbers have the same length
	// Map to structs
	var races []race
	for i := 0; i < len(timeNumbers); i++ {
		races = append(races, race{
			time:     timeNumbers[i],
			distance: distanceNumbers[i],
		})
	}

	multiply := 1
	for _, race := range races {

		count := 0
		for i := 0; i <= race.time; i++ {

			speed := i
			time := race.time - i

			distance := speed * time

			if distance > race.distance {
				count++
			}

		}

		//fmt.Println(count)

		multiply = multiply * count
	}

	return multiply
}

// Function to extract numbers from a string
func extractNumbers(s string) []int {
	var numbers []int
	fields := strings.Fields(s)
	for _, field := range fields {
		if num, err := strconv.Atoi(field); err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func extractNumbersPartTwo(s string) []int {
	var numbers []int
	var numString string
	fields := strings.Fields(s)
	for i := 1; i < len(fields); i++ {
		numString += fields[i]
	}

	if num, err := strconv.Atoi(numString); err == nil {
		numbers = append(numbers, num)
	}
	return numbers
}

func partTwo(input string) int {

	file, _ := os.ReadFile(input)
	lines := strings.Split(string(file), "\n")

	// Extract numbers from each line
	timeNumbers := extractNumbersPartTwo(lines[0])
	distanceNumbers := extractNumbersPartTwo(lines[1])

	// Check if timeNumbers and distanceNumbers have the same length
	// Map to structs
	var races []race
	for i := 0; i < len(timeNumbers); i++ {
		races = append(races, race{
			time:     timeNumbers[i],
			distance: distanceNumbers[i],
		})
	}

	fmt.Println(races)

	multiply := 1
	for _, race := range races {

		count := 0
		for i := 0; i <= race.time; i++ {

			speed := i
			time := race.time - i

			distance := speed * time

			if distance > race.distance {
				count++
			}

		}

		//fmt.Println(count)

		multiply = multiply * count
	}

	return multiply
}
