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
