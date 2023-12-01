package modules

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PartOne() {
	var sum int = 0

	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var min string
		var max string

		line := scanner.Text()

		array := []byte(line)

		for _, c := range array {
			if c >= '0' && c <= '9' {
				min = string(c)
				break
			}
		}

		for i := len(array) - 1; i >= 0; i-- {
			if array[i] >= '0' && array[i] <= '9' {
				max = string(array[i])
				break
			}
		}

		lineInt, _ := strconv.Atoi(min + max)

		// fmt.Println(line, min, max)

		sum += lineInt

	}

	fmt.Println(sum)
}
