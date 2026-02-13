//
//A palindromic number reads the same both ways. The largest palindrome made from the product of two $2$-digit numbers is 9009 = 91 * 99
//Find the largest palindrome made from the product of two 3-digit numbers.
// solution: this one is a bit inellegant but we get there. we first write a fxn that detects palindrome numbers by reversing them and checking.
// we loop over all the 3 digits numbers i, for each i we loop again over all 3-digit numbers j <= i (we don't want make a full times table, only the upper half, saves time and memory)
// we first check if the product of i and j are bigger than our current candidate palindrome, if it is, we check if it's a palindrome, if it is we assign it to largest palindrome.
// once we exit the outer loop we're done and we just print the largest paindrome.
// This approach is basically the naive approah where we add a memory and time optimization by only calcuating products we haven't already calculated and we don't store a slice
// of all the known palindromes but instead only store the biggest one.
// It's pretty brutish but it's OK for now.

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
