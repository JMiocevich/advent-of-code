package main

import "testing"

func TestPartOne(t *testing.T) {

	want := 114

	res := partOne("testinput")

	if want != res {
		t.Fatalf("Got %d wanted %d", res, want)
	}

}

func TestPartTwo(t *testing.T) {

	want := 5

	res := partTwo("testinput2")

	if want != res {
		t.Fatalf("Got %d wanted %d", res, want)
	}

}
