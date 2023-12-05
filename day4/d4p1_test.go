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

	res := getWinningValue("test.txt")

	fmt.Println("Output : ", res)

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

		// fmt.Println(winningCards)
		// fmt.Println(cards)
		// put myCards into map
		cardMap := make(map[string]bool)

		for _, card := range cards {
			if card != "" {
				cardMap[card] = true
			}
		}
		// check if number is in winning deck

		for _, card := range winningCards {

			if cardMap[card] {
				power += 1

			}
		}

		if power != -1 {
			sum += math.Pow(float64(2), float64(power))
		}

	}

	return sum
}
