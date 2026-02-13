//
//A palindromic number reads the same both ways. The largest palindrome made from the product of two $2$-digit numbers is 9009 = 91 * 99
//Find the largest palindrome made from the product of two 3-digit numbers.

package main

import "fmt"

func main() {
	largestPal := 0
	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			if i*j > largestPal && isPalindrome(i*j) {
				largestPal = i * j
			}
		}
	}
	fmt.Println("the largest palindrome that is the procudct of 2 3-digit number is", largestPal)
}

func isPalindrome(n int) bool {
	num := n
	rev := 0
	for num > 0 {
		dig := num % 10
		rev = rev*10 + dig
		num = num / 10
	}
	if n == rev {
		return true
	} else {
		return false
	}
}
