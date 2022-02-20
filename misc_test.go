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
