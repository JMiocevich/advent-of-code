package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	res := partOne("input.txt")
	fmt.Println(res)

}

func partOne(fileInput string) int {

	file, _ := os.Open(fileInput)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {

		line := scanner.Text()

		gameNumber := strings.Split(line, ":")

		games := strings.Split(gameNumber[1], ";")

		red := 0
		green := 0
		blue := 0

		for _, game := range games {

			pairs := strings.Split(game, ",")

			for _, pair := range pairs {

				pair = strings.TrimSpace(pair)

				part := strings.Split(pair, " ")

				// fmt.Println(part[0])
				// fmt.Println(string(part[1]))

				v, _ := strconv.Atoi(part[0])

				if string(part[1]) == "red" && v > red {

					red = v
				}
				if string(part[1]) == "green" && v > green {
					green = v
				}
				if string(part[1]) == "blue" && v > blue {
					blue = v
				}

			}
		}

		power := red * green * blue

		sum += power

		// break
	}

	return sum
}
