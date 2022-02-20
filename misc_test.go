package main

import "testing"

func TestNilSlices(t *testing.T) {
	SliceWithNilCheck(nil)
	SliceNoNilCheck(nil)

	emptySlice := []string{}
	SliceWithNilCheck(emptySlice)
	SliceNoNilCheck(emptySlice)

	simpleSlice := []string{"alpha", "beta"}
	SliceWithNilCheck(simpleSlice)
	SliceNoNilCheck(simpleSlice)
}

func TestValueSwap(t *testing.T) {
	a, b := "alpha", "beta"
	SwapValuesWithTemp(&a, &b)
	if a != "beta" || b != "alpha" {
		t.Fail()
	}
	SwapValuesWithTemp(&a, &b)
	if a != "alpha" || b != "beta" {
		t.Fail()
	}
}
