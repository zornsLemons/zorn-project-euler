// problem: Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
// solution: becasue of Gausses formulas you can do this on a pocket calculator but we can do it in Go for fun.
// the sum of the first n naural numbers is (n)(n-1)/2 the sum of the squares of the first n natural numbers is (n)*(n+1)(2n+1)/6
package main

import (
	"fmt"
)

func main() {
	n := 100

	naturalsSum := ((n) * (n + 1)) / 2
	naturals2Sum := ((n) * (n + 1) * (2*n + 1)) / 6

	var diff int

	if (naturalsSum * naturalsSum) > naturals2Sum {
		diff = ((naturalsSum * naturalsSum) - naturals2Sum)
	} else {
		diff = (naturals2Sum) - (naturalsSum * naturalsSum)
	}

	fmt.Println("the difference in the sum of squares and the squared sum of teh first 100 natural numbers is", diff)
}
