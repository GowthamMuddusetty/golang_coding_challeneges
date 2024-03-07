package codingchallenege100days

import (
	"fmt"
	"math"
)

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
