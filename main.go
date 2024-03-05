package main

import "fmt"

// isPrime checks if a number is prime.
func isPrime(n int) bool {
	c := make(chan int)
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		go calcPrime(n, i, c)
		if <-c == 0 {
			return false
		}
	}
	return true
}

func calcPrime(n int, i int, c chan<- int) {
	a := n % i
	c <- a
}

// generatePrimes generates prime numbers up to a given limit.
func generatePrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// isPrime checks if a number is prime.
func isPrimeOld(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// generatePrimes generates prime numbers up to a given limit.
func generatePrimesOld(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrimeOld(i) {
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
