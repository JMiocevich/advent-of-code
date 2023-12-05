package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

// func TestPartTwo(t *testing.T) {

// 	res := getWinningValueP2("test_input.txt")

// 	if int(res) != 30 {
// 		t.Fatalf("Expected 30, got %d", int(res))
// 	}

// 	fmt.Println(res)

// }

func TestPartTwoReal(t *testing.T) {

	res := getWinningValueP2("input.txt")

	fmt.Println(int(res))

}

func getWinningValueP2(input string) float64 {
	file, _ := os.Open(input)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0.0

	var arrayLen [250]int
	arrPos := 0

	for scanner.Scan() {
		arrayLen[arrPos] += 1
		count := 0
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
				count += 1
			}
		}
		fmt.Println("arrPos", arrPos)
		for i := 0; i < count; i++ {
			arrayLen[arrPos+i+1] += 1 * arrayLen[arrPos]

			fmt.Println("arrayLen", arrayLen)
		}

		arrPos++

	}

	fmt.Println(arrayLen)

	for _, val := range arrayLen {
		sum += float64(val)
	}
	return sum
}
