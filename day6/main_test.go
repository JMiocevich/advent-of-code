package main

import "testing"

func TestPartOne(t *testing.T) {

	want := 288

	res := partOne("testinput")

	if res != want {
		t.Errorf("got %d, wanted %d", res, want)
	}

}

func TestPartTwo(t *testing.T) {

	want := 71503

	res := partTwo("testinput")

	if res != want {
		t.Errorf("got %d, wanted %d", res, want)
	}

}
