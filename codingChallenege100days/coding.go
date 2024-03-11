package codingchallenege100days

import (
	"fmt"
	"math"
)

// 1.  Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0."
func Reverse(x int) int {
	reversed := 0
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	for x > 0 {
		digit := x % 10
		if reversed > math.MaxInt32/10 || (reversed == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		reversed = reversed*10 + digit
		x /= 10
	}
	return reversed * sign
}

// 2.  Write a program that prints the numbers from 1 to 100. But for multiples of three, print 'Fizz' instead of the number,
// and for the multiples of five, print 'Buzz'. For numbers that are multiples of both three and five, print 'FizzBuzz'.")
func Challenge1() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz", ",")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuz")
		}
	}
}

// 3.FizzBuzz  The task is to write a program that prints the numbers from 1 to 100.
// But for multiples of three, print "Fizz" instead of the number, and for the multiples of five, print "Buzz".
//
//	For numbers that are multiples of both three and five, print "FizzBuzz".
func FizzBuzz() {

	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz-- ", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz-- ", i)
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz-- ", i)
		}
	}
}

// 4. Reverse a string
func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	for i := 0; i < len(runes); i++ {
		fmt.Println(runes, " ", i)
	}

	return string(runes)
}

// 5. Palindrome Check
func PalindromeCheck(str string) bool {

	// 1 way

	// revStr := ""
	// runes := []rune(str)
	// for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
	// 	runes[i], runes[j] = runes[j], runes[i]
	// }
	// revStr = string(runes)
	// if str == revStr {
	// 	fmt.Println("is Palindrome")
	// } else {
	// 	fmt.Println("is not Palindrome")
	// }

	// 2nd way
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

// 6. Fibanocci Numbers
func FibanocciNumbers(num int) {

	// 1 way
	// fib := make([]int, num)
	// fib[0] = 0
	// fib[1] = 1

	// for i := 2; i < num; i++ {
	// 	fib[i] = fib[i-1] + fib[i-2]
	// }
	// return fib

	// 2nd way
	num1 := 0
	num2 := 1
	num3 := 0
	fmt.Print(num1, num2, " ")

	for i := 2; i < num; i++ {
		num3 = num1 + num2
		num1 = num2
		num2 = num3
		fmt.Print(num3, " ")
	}
}

// 7. Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// you can return the answer in any order.
func TwoSums(nums []int, target int) []int {
	numsMap := make(map[int]interface{})

	for i := 0; i < len(nums); i++ {
		reqNum := target - nums[i]
		reqNumIdx, isPresent := numsMap[reqNum]

		if isPresent {
			return []int{i, reqNumIdx.(int)}
		}
		numsMap[nums[i]] = i

	}
	// if it doesnt found any such pair it will return this numbers
	return []int{322, 232}
}
