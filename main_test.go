package main

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{1, 0},
		{2, 1},
		{5, 3},
		{10, 4},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			primes := generatePrimes(tt.a)
			if len(primes) != tt.want {
				t.Errorf("got %v, want %v", len(primes), tt.want)
			}
		})
	}
}

func BenchmarkGeneratePrimesNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generatePrimes(100)
	}
}

func BenchmarkGeneratePrimesOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generatePrimesOld(100)
	}
}
