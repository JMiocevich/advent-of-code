package main

import "testing"

type mapType struct {

    destination int
    source int
    offSet int
}

func TestPartOne(t *testing.T) {

    want := 35




    location := 1
    
   // location -> humidity
    getHumdity(location, htol) 

}


func getHumdity(location int ,htol []mapType) int {
    
    for _, map := range htol {

	delta := map.destination - map.source
	
	if location >= map.source && location < map.source + map.offSet {
	    
	    source := location - delta
	}else {
	    source := location
	}

    }


    return 0
}
