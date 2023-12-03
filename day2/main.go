package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	res := partOne("input.txt")
	fmt.Println(res)

}

func partOne(fileInput string) int {

	file, _ := os.Open(fileInput)
	defer file.Close()
 	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		fmt.Println(line)
		
	}

	return 6
}
