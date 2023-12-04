package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {

	want := 13.0
	res := getWinningValue("test_input.txt")

	if res != want {
		t.Fatalf(`Its not working, value returned: %f`, res)
	}

	fmt.Println(res)

}

func TestPartOneReal(t *testing.T) {

	res := getWinningValue("input.txt")

	fmt.Println(res)

}

func getWinningValue(input string) float64 {
	file, _ := os.Open(input)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0.0

	for scanner.Scan() {
		power := -1
		line := scanner.Text()

		allCards := strings.Split(line, ":")[1]

		winningCards := strings.Split(strings.TrimSpace(strings.Split(allCards, "|")[0]), " ")
		cards := strings.Split(strings.TrimSpace(strings.Split(allCards, "|")[1]), " ")

		// put myCards into map
		cardMap := make(map[string]bool)

		for _, card := range cards {
			cardMap[card] = true
		}

		// check if number is in winning deck

		for _, card := range winningCards {

			if cardMap[card] {
				power += 1
			}
		}

		if power == -1 {
			continue
		}

		sum += math.Pow(float64(2), float64(power-1))

	}

	return sum
}
