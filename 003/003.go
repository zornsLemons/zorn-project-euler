// The prime factors of 13195 are 5, 7, 13 and 29
// What is the largest prime factor of the number 600851475143?
// Solution: Prime facotrization is at the center of RSA for a damn REASON. Not a lot to do here except do it the hard way.
// I generate all primes up to sqrt(N) using the sieve of eratosthenes the sieve algo gives me an ordered list of primes and Since
// I only want the largest prime divisor of N and don't care about the other factors, I loop backward through the slice of primes and
// simply print then first one that evenly divides N.
// this will eat all available memory on the planet for large N.

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// testing
	start := time.Now()
	N := 600851475143
	K := int(math.Ceil(math.Sqrt(float64(N))))
	candidatePrimes := GenPrimesBelowK(K)
	for i := len(candidatePrimes) - 1; i >= 0; i-- {
		if N%candidatePrimes[i] == 0 {
			fmt.Println("largest factor is:", candidatePrimes[i])
			break
		}
	}
	duration := time.Since(start)
	fmt.Println("ran in: ", duration)
}

func GenPrimesBelowK(K int) []int {
	primes := make([]bool, K+1)
	for i := 2; i <= K; i++ {
		primes[i] = true
	}
	for p := 2; p*p < K; p++ {
		if primes[p] {
			for j := p * p; j <= K; j += p {
				primes[j] = false
			}
		}
	}
	var primeNumbers []int
	for p, val := range primes {
		if val {
			primeNumbers = append(primeNumbers, p)
		}
	}
	return primeNumbers
}
