package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// isPalindrome checks whether the given string is a palindrome.
func isPalindrome(s string) bool {
	// Remove spaces and convert to lowercase
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	// Check if the string reads the same forward and backward
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Palindrome Checker")
	fmt.Println("===================")
	fmt.Println("Enter a string to check if it is a palindrome (type 'exit' to quit):")

	for {
		fmt.Print("Enter a string: ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if isPalindrome(input) {
			fmt.Printf("\"%s\" is a palindrome!\n", input)
		} else {
			fmt.Printf("\"%s\" is not a palindrome.\n", input)
		}
	}
}
