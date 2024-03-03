package main

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	var tests = []struct {
		a    int
		want bool
	}{
		{-1, false},
		{1, false},
		{2, true},
		{16, false},
		{17, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		p := make(chan bool)
		t.Run(testname, func(t *testing.T) {
			go isPrime(tt.a, p)
			if <-p != tt.want {
				t.Errorf("got %v, want %v", <-p, tt.want)
			}
		})
	}
}

func BenchmarkGeneratePrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go generatePrimes(1000)
	}
}
