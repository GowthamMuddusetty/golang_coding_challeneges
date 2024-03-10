package main

import (
	code "codingQuestions/codingChallenege100days"
	"fmt"
)

func main() {

	// 1.  Given a signed 32-bit integer x, return x with its digits reversed.
	// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.")
	fmt.Println(code.Reverse(123))

	// 2.  Write a program that prints the numbers from 1 to 100. But for multiples of three, print 'Fizz' instead of the number,
	// and for the multiples of five, print 'Buzz'. For numbers that are multiples of both three and five, print 'FizzBuzz'.")
	code.Challenge1()

	// 3.FizzBuzz  The task is to write a program that prints the numbers from 1 to 100.
	// But for multiples of three, print "Fizz" instead of the number, and for the multiples of five, print "Buzz".
	//  For numbers that are multiples of both three and five, print "FizzBuzz".
	code.FizzBuzz()

	// 4. Reverse a string
	fmt.Println(code.ReverseString("hello world"))

	// 5. Palindrome Check
	bool := code.PalindromeCheck("racecar")
	if bool {
		fmt.Println("Is Palindrome")
	} else {
		fmt.Println("Is Not Palindrome")
	}
}
