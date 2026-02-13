// The prime factors of 13195 are 5, 7, 13 and 29
// What is the largest prime factor of the number 600851475143?
// Solution: Prime facotrization is at the center of RSA for a damn REASON. Not a lot to do here except do it the hard way.
// We could do this by reading from a text file that we populate with prime numbers up to n/2 but like...that feels like cheating and
// not really what I want ot do even though it's easier. Instead, I'm going to just the test division method and see if I can come up with
// any cute ways to cut down the search space as we go.
//
// We're going use a binary tree data structure.
package main

func main() {
}

// this is our factor tree data structure.
type Node struct {
	value uint
	left  *Node
	right *Node
}

func factor(number uint) (uint, uint) {
}
