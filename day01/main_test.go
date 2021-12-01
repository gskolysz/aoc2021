package main

import "testing"

func TestSumInts(t *testing.T) {
	ints := []int{1, 2, 3}
	total := SumInts(ints)
	if total != 6 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 6)
	}
}

func TestCountWindowedIncreases(t *testing.T) {
	ints := []int{607, 618, 618, 617, 647, 716, 769, 792}
	total := CountWindowedIncreases(ints)
	if total != 5 {
		t.Errorf("CountWindowedIncreases was incorrect, got: %d, want: %d.", total, 5)
	}
}
