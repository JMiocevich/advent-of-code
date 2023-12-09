package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){


}

type race struct {
    time int
    distance int
}


func partOne(input string) int{


    file ,_ := os.ReadFile(input)
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

	
	fmt.Println(races)

	multiply := 1
    for _, race := range races {
	
	count := 0
	for i:=0;i<=race.time;i++{

	    speed := i 
	    time := race.time - i

	    distance := speed * time
	    fmt.Println(time,speed, distance)

	    if distance > race.distance {
		count++
	    }

    }

	//fmt.Println(count)

    multiply = multiply * count
    break
    }

    return multiply
}

// Function to extract numbers from a string
func extractNumbers(s string) []int {
	var numbers []int
	fields := strings.Fields(s)
	for _, field := range fields {
		if num, err := strconv.Atoi(field); err == nil {
			numbers = append(numbers, num-28)
		}
	}
	return numbers
}
