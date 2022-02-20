package main

import (
	"log"
)

// Iterates over a slice with a redundant nil check.
func SliceWithNilCheck(slice []string) {
	// nolint:gosimple // for demonstration purposes
	if slice != nil {
		for _, s := range slice {
			log.Println(s)
		}
	}
}

// Iterates over a slice without a nil check. Unlike other
// languages - such as Java - Go does not require a nil
// check when iterating over an empty or nil slice.
func SliceNoNilCheck(slice []string) {
	for _, s := range slice {
		log.Println(s)
	}
}

// Swaps values of the provided pointers using
// an unnecessary temporary variable.
func SwapValuesWithTemp(a *string, b *string) {
	// nolint:gocritic // for demonstration purposes
	temp := *a
	*a = *b
	*b = temp
}

// Swaps values of the provided pointers without a temporary variable.
func SwapValuesNoTemp(a *string, b *string) {
	*a, *b = *b, *a
}
