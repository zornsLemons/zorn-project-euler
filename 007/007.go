// what is the 10,001st prime number?
//
// Solution: the obvious and stupid thing to do is to just check every number, store the primes in a list as you
// find them by dividing by all numbers smaller than the test number, and then stop when your prime list reaches len 10,001.
// We're not doing that. Instead, we're going to note that it suffices to check that the primes you already know for divisibility
// b/c of the fundemental thm of arithematic.
// second, we only need to check primes less than the sqrt of the check value.
// third, we only need to check odd numbers.
// fourth, we can run the divisibility checks concurrently to get a speed up.
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	const primeIdx = 100000
	var primes []int
	primes = append(primes, 3)
	for i := 2; ; i++ {
		num := 2*i + 1
		isPrime := isPrimeWithPrimeList(num, primes)
		if len(primes) == primeIdx {
			break
		}
		if isPrime {
			primes = append(primes, num)
			continue
		}
		continue
	}
	run := time.Since(start)
	fmt.Println("The 10,001st prime is:", primes[primeIdx-1])
	fmt.Println("calculated in ", run)
}

// func maxIdxLessThanNumSq(num int, numSlice []int) int {
// 	for i, val := range numSlice {
// 		if val*val < num {
// 			return i
// 		}
// 		continue
// }
// 	return len(numSlice) - 1
// }

func isPrimeWithPrimeList(num int, primes []int) bool {
	// maxID := maxIdxLessThanNumSq(num, primes)
	// maxDiv := num * num
	for _, prime := range primes {
		if num%prime == 0 {
			return false
		}
		continue
	}
	return true
}

// func isPrimeWithPrimeList(num int, primes []int) bool {
// 	maxIdx := maxIdxLessThanNumSq(num, primes)
// 	divQ := make(chan int, maxIdx)
// 	done := make(chan bool, 1)
// 	composite := false
// 	for i := 0; i <= runtime.NumCPU(); i++ {
// 		go func() {
// 			checkDiv, more := <-divQ
// 			if more {
// 				if num%checkDiv == 0 {
// 					composite = true
// 					done <- true
// 					return
// 				}
// 			} else {
// 				done <- true
// 				return
// 			}
// 		}()
// 	}
// 	for _, divCand := range primes {
// 		divQ <- divCand
// 	}
// 	close(divQ)
// 	<-done
// 	return composite
// }
