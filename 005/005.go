// problem
// 2520 is the smallest number that can be divided by each of the numbers from $1$ to $10$ without any remainder.
// What is the smallest positive number that is evenly divisible (divisible with no remainder) by all of the numbers from $1$ to $20$?
// solution: Let's let N be the upper bound on our range (20 in this case) and x be the unknown number that all numbers <N evenly divides.
// The largest prime power < N will be a factor of x for every prime <N. this gives us a number for which any number less than N can be constructed from
// a subset of x's prime factorization. this number is minimal almost by force.
// this example is probably easier to work out by hand BUT for large N not so much so we write the algo without magic numbers.
// init the result as 1.
// The algo is straight forward, generate all primes <N, for each prime take succsesive powers until the calue exceeds N, at each step multiply
// x by the prime. Do this until we're through all the primes. x is now the result we seek.

package main

import (
	"fmt"

	"github.com/zornsLemons/zorn-project-euler/eulerutils"
)

func main() {
	N := 20
	primes := eulerutils.GenPrimesBelowK(N)
	x := 1
	for _, p := range primes {
		for val := p; val < N; val = val * p {
			x = x * p
		}
	}
	fmt.Println("the smallest number divisible by all numbers between 1 and 20 is:", x)
}
