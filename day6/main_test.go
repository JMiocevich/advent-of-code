package main 

import"testing"

func TestPartOne(t *testing.T){

    want := 288

    res := partOne("input")
    
    if res != want {
    t.Errorf("got %d, wanted %d" , res ,want) }


}

