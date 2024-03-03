package main

import "fmt"

// isPrime checks if a number is prime.
func isPrime(n int, p chan bool) {
	go func() {
		defer close(p)
		if n <= 1 {
			p <- false
		}

		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				p <- false
			}
		}
		p <- true
	}()
}

func generatePrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		p := make(chan bool)
		go isPrime(i, p)
		if <-p {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	limit := 100
	primes := generatePrimes(limit)
	fmt.Println(primes)
}
