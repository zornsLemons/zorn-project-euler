// Package eulerutils
// some utilities for project euler problems
package eulerutils

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
