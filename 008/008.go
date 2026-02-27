// Problem 8: The four adjacent digits in the $1000$-digit number that have the greatest product are $9 \times 9 \times 8 \times 9 = 5832$.
//
// <p class="monospace center">
// 73167176531330624919225119674426574742355349194934<br>
// 96983520312774506326239578318016984801869478851843<br>
// 85861560789112949495459501737958331952853208805511<br>
// 12540698747158523863050715693290963295227443043557<br>
// 66896648950445244523161731856403098711121722383113<br>
// 62229893423380308135336276614282806444486645238749<br>
// 30358907296290491560440772390713810515859307960866<br>
// 70172427121883998797908792274921901699720888093776<br>
// 65727333001053367881220235421809751254540594752243<br>
// 52584907711670556013604839586446706324415722155397<br>
// 53697817977846174064955149290862569321978468622482<br>
// 83972241375657056057490261407972968652414535100474<br>
// 82166370484403199890008895243450658541227588666881<br>
// 16427171479924442928230863465674813919123162824586<br>
// 17866458359124566529476545682848912883142607690042<br>
// 24219022671055626321111109370544217506941658960408<br>
// 07198403850962455444362981230987879927244284909188<br>
// 84580156166097919133875499200524063689912560717606<br>
// 05886116467109405077541002256983155200055935729725<br>
// 71636269561882670428252483600823257530420752963450>
// Find the thirteen adjacent digits in the $1000$-digit number that have the greatest product. What is the value of this product?
//
// Solution: We're going to read the big number in from a file.
// We store the number in memory as a slice of its digits.
// we dont need to store the sub slices or  products of ubslices in memory.
// instead we need to only store the biggest one so far as we loop over the big int slice.
//
//	the first code was slow becasue I'm calculating the full array product at each step. I make this fast by only
//	calculating array products when the leading digit is a 0 and when it isn't I get the next product by
//	dividng the last product by first element of it's subsequence and multiplying by the first element outside
//	the sub sequence.
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := "number.txt"

	bigInt := bigNumberHandler(filepath)
	fmt.Println(len(bigInt))
	fmt.Println(bigInt[:13])
	biggestProduct := arrayProduct(bigInt[0:13])
	fmt.Println(len(bigInt[0:13]))
	// candidate := biggestProduct
	var candidate int
	// var bigSlice []int
	for i := 0; i+15 < len(bigInt); i++ {
		switch bigInt[i] {
		case 0:
			candidate = arrayProduct(bigInt[i+1 : i+14])
		default:
			candidate = (candidate * bigInt[i+13]) / bigInt[i]
		}
		// candidate = arrayProduct(bigInt[i+1 : i+14])
		if candidate >= biggestProduct {
			biggestProduct = candidate
		}
	}
	fmt.Println("The biggest product is ", biggestProduct)
	// fmt.Println("the big subseq is ", bigSlice)

	// fmt.Println(len(bigSlice))
}

func bigNumberHandler(filePath string) []int {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Fatal Error reading File")
		log.Fatal(err)
	}
	strContent := string(content)
	return digitStringToDigitSlice(strContent)
}

func digitStringToDigitSlice(string string) []int {
	runeSlice := []rune(string)
	digSlice := make([]int, len(runeSlice))
	for i, r := range runeSlice {
		if r >= '0' && r <= '9' {
			digSlice[i] = int(r - '0')
		}
	}
	return digSlice
}

func arrayProduct(intArray []int) int {
	runningProduct := 1
	for _, num := range intArray {
		runningProduct = num * runningProduct
	}
	return runningProduct
}
