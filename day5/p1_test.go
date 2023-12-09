package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

type mapType struct {
	destination int
	source      int
	offSet      int
}

func TestPartOne(t *testing.T) {

	file, _ := os.ReadFile("input")
	var arrayOfMaps [][]mapType

	seedLine := strings.Split(string(file), "\n")[:1]
	seedValues := strings.Split(strings.TrimSpace(strings.Split(seedLine[0], ":")[1]), " ")

	var seeds []int
	for _, s := range seedValues {
		seed, _ := strconv.Atoi(s)
		seeds = append(seeds, seed)
	}

	maps := strings.Split(string(file), "\n\n")[1:]

	for _, m := range maps {
		lines := strings.Split(m, "\n")
		var mapArray []mapType

		for _, line := range lines[1:] {
			values := strings.Fields(line)
			if len(values) == 3 {
				dest, _ := strconv.Atoi(values[0])
				src, _ := strconv.Atoi(values[1])
				offSet, _ := strconv.Atoi(values[2])
				mapArray = append(mapArray, mapType{destination: dest, source: src, offSet: offSet})
			}
		}

		arrayOfMaps = append(arrayOfMaps, mapArray)
	}

	right := 9999999999
	var pointer int

	for z := 1; z < right; z++ {

		pointer = z

		for i := 6; i >= 0; i-- {
			pointer = getSource(pointer, arrayOfMaps[i])
		}

		if contains(seeds, pointer) {
			fmt.Printf("Lowest Location is : %d \n", z)
			break
		}
	}

}

func contains(slice []int, value int) bool {

	for i := 0; i < len(slice)-1; i = i + 2 {
		// fmt.Println(value, slice[i], slice[i+1])
		if value >= slice[i] && value < slice[i]+slice[i+1] {
			return true
		}

	}
	return false
}

func getSource(destination int, Map []mapType) int {

	for _, section := range Map {

		if destination >= section.destination && destination < section.destination+section.offSet {

			delta := (destination - section.destination) + section.source

			return delta

		}
	}
	return destination

}
