package main

import (
	linkedin "codingQuestions/LinkedinCodingQs"
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

	// 6. Fibanocci Numbers
	code.FibanocciNumbers(10)

	// 7. Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
	// You may assume that each input would have exactly one solution, and you may not use the same element twice.
	// you can return the answer in any order.
	nums := []int{2, 11, 7, 15}
	//nums := []int{21, 11, 17, 15}
	indexes := code.TwoSums(nums, 9)
	fmt.Println(indexes)

	// Linkedin Questions

	// 1. Problem: Generating 4-digit random number
	linkedin.GenRandom4Digit()

	// 2. Problem: Identify the count of a letter in the sentence
	occurences := linkedin.CountLetterOccurence("Gods and Heroes")
	fmt.Println(occurences)

}
