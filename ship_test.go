package main

import (
	"fmt"
	"math"
	"testing"
)

func TestMod360(t *testing.T) {
	tests := []struct {
		N        float64
		Expected float64
	}{
		{N: 0, Expected: 0},
		{N: 10, Expected: 10},
		{N: 90, Expected: 90},
		{N: 360, Expected: 0},
		{N: 365, Expected: 5},
		{N: 360 + 90, Expected: 90},
		{N: -10, Expected: 360 - 10},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.N), func(tt *testing.T) {
			if mod360(test.N) != test.Expected {
				t.Errorf("expected to get %v for %v but got %v", test.Expected, test.N, mod360(test.N))
			}
		})
	}
}

func TestHeadings(t *testing.T) {
	deg := 0.0
	rad := deg * math.Pi / 180.0
	fmt.Println(math.Sincos(rad))
}
