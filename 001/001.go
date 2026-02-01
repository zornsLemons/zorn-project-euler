// Problem 1:
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3,5,6,9 and .
// The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
// Solution Note: it is no more difficult to find the soltuion for all multiples of 3 and 5
// <N for large(ish) N. So we solve that instead.
// This solution is very simple, iterate over numbers less than 1000, if 3 or 5 divides it, we add it to a runing sum.

package main

import "fmt"

func main() {
	soltuion := sumMultiples(1000)
	fmt.Println("The sum of the multiples of 3 or 5 less than 1000 is: ", soltuion)
}

func hasFac3o5(num int) int {
	if num%3 == 0 || num%5 == 0 {
		return num
	}
	return 0
}

func sumMultiples(rangeMax int) int {
	sum := 0
	for i := 1; i < rangeMax; i++ {
		sum = sum + hasFac3o5(i)
	}

	return sum
}
