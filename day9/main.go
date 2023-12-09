package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	res := partOne("input")

	fmt.Println("ANSWER: ", res)

	res2 := partTwo("input")

	fmt.Println("ANSWER Part 2: ", res2)

}

func partOne(input string) int {

	file, _ := os.Open(input)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lastValues []int

	for scanner.Scan() {

		numbers := extractNumbers(scanner.Text())

		value := recurse(numbers)

		lastValues = append(lastValues, value)

	}

	return sum(lastValues)

}

func partTwo(input string) int {

	file, _ := os.Open(input)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lastValues []int

	for scanner.Scan() {

		numbers := extractNumbers(scanner.Text())

		value := recurse2(numbers)

		lastValues = append(lastValues, value)

	}

	return sum(lastValues)
}

func recurse2(array []int) int {

	if arrayZero(array) {
		return array[0]
	}

	nextArray := calculateNext2(array)
	firstValue := recurse2(nextArray)

	return firstValue + array[0]

}

func recurse(array []int) int {

	if arrayZero(array) {
		return array[len(array)-1]
	}

	nextArray := calculateNext(array)

	lastValue := recurse(nextArray)

	return lastValue + array[len(array)-1]

}

func sum(values []int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func calculateNext2(array []int) []int {

	var newArray []int

	for i := 0; i < len(array)-1; i++ {
		newArray = append(newArray, array[i]-array[i+1])
	}
	return newArray
}

func calculateNext(array []int) []int {

	var newArray []int

	for i := 0; i < len(array)-1; i++ {
		newArray = append(newArray, array[i+1]-array[i])
	}
	return newArray
}

func arrayZero(array []int) bool {

	for _, field := range array {
		if field != 0 {
			return false
		}
	}
	return true
}

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
