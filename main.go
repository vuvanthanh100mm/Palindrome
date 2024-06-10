
package main

import (
	"fmt"
	"strconv"
)

// isPalindrome checks if a given positive integer n is a palindrome
func isPalindrome(n int) bool {
	// Convert the integer to a string
	str := strconv.Itoa(n)
	
	// Get the length of the string
	length := len(str)
	
	// Compare characters from both ends
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	testNumbers := []int{121, 12321, 123, 1, 0, 1221, 123456}
	
	for _, number := range testNumbers {
		fmt.Printf("Is %d a palindrome? %v\n", number, isPalindrome(number))
	}
}