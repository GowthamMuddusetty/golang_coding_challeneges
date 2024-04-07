package linkedin

import (
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

/* 1. Problem: Generating 4-digit random number */
func GenRandom4Digit() {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 10000
	randNum := rand.Intn(max-min) + min
	if randNum >= 1000 && randNum <= 9999 {
		fmt.Println(randNum)
	}
}

/* 2. Problem: Identify the count of a letter in the sentence */
func CountLetterOccurence(str string) string {
	runes := []rune(strings.ToLower(str))
	strMap := make(map[string]int)
	for i := 0; i < len(runes); i++ {
		char := string(runes[i])

		/*
			In Go, strMap[char]++ is a shorthand for incrementing the value associated with the key char in the strMap by 1.

			If the key char doesn't exist in the map, it will be initialized with the zero value for its type, which is 0 for integers.
			If it already exists, it will be incremented by 1.This is a concise way to update the count of occurrences for each character in the string.

			val := strMap[char]
			strMap[char] = val + 1     			// this is same as strMap[char]++
		*/
		strMap[char]++
	}

	var result string
	for char, count := range strMap {
		result += fmt.Sprintf("%s:%d ", char, count)
	}

	return result
}

/* 4. Generate random without using math.rand function */
func GenRandom4DigitNoRand() {
	rand.Seed(time.Now().UnixNano())

	// generate a random 8-digit ID
	id := ""
	for i := 0; i < 8; i++ {
		id += strconv.Itoa(rand.Intn(10))
	}
}

/* 5. Problem: Finding number in an Array and ArrayList */
func Find_Num() {
	var nums []int
	nums = append(nums, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fNums := []int{3, 7, 11}
	fmt.Println(nums)
	for i := 0; i < len(fNums); i = i + 1 {
		for j := 0; j < len(nums); j = j + 1 {
			if fNums[i] == nums[j] {
				fmt.Println(fNums[i], "exists in nums")
			}
		}
	}
}

/* 6. Problem: Find out highest numbers in an Array whose adjacent right numbers are smaller */
func Highest_Num() {
	var nums [6]int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 4
	nums[3] = 3
	nums[4] = 5
	nums[5] = 2

	for i := 0; i < len(nums)-1; i = i + 1 {
		if nums[i] > nums[i+1] {
			fmt.Print(nums[i], ",")
		}
	}
	if nums[len(nums)-1] > nums[0] {
		fmt.Print(nums[len(nums)-1])
	}
}

/*7. Problem: Find a count of each character in a given String.*/
func Char_Count() {
	str := "Automation Testing"
	fmt.Println(len(str))

	runes := []rune(strings.ToLower(str))
	fmt.Println(len(runes))

	chrCount := make(map[string]int)
	for i := 0; i < len(runes); i++ {
		char := string(runes[i])
		chrCount[char]++
	}

	var result string
	for char, count := range chrCount {
		result += fmt.Sprintf("%s:%d,", char, count)
	}
	fmt.Println(chrCount)
	fmt.Println(result)
}

/* 8. Problem: Extract last 4 characters of any String.*/
func Extract_Char() {
	str := "Backend Development"
	newStr := strings.ToLower(str)
	str = strings.ReplaceAll(newStr, " ", "")

	//result := ""
	// for _, v := range str {
	// 	result = string(v) + result
	// 	// fmt.Println(string(v))
	// 	// fmt.Println(result)
	// }

	// for i := 0; i < 4; i++ {
	// 	fmt.Print(string(result[i]) + "\n")
	// }

	// 2nd way
	runes := []rune(str)
	for i := len(runes) - 1; i > len(runes)-5; i-- {
		fmt.Println(string(runes[i]))
	}
}

/* 9. Problem: Program to remove duplicates from a given array. */
func RemDup() {
	var nums = [7]int{2, 3, 5, 5, 8, 2, 3}
	var dup = []int{}
	for i := 0; i < len(nums); i++ {
		if !slices.Contains(dup, nums[i]) {
			dup = append(dup, nums[i])
		}
	}
	fmt.Println(dup)
}

/* 10. Problem: First letter of each word present in String. */
func FirstLet() {
	str := "It's not the load that breaks you down, it's the way you carry it."
	words := strings.Fields(str)
	var arrStr []string
	for _, word := range words {
		chrs := string(word[0])
		arrStr = append(arrStr, chrs)
	}
	fmt.Println(strings.Join(arrStr, ","))
}

/* 11. Problem: Fibonacci Series using Arrays. */
func FibSeries() {
	num := 10
	fibSer := make([]int, num)
	fibSer[0] = 0
	fibSer[1] = 1
	for i := 2; i < num; i++ {
		sum := fibSer[i-2] + fibSer[i-1]
		fibSer[i] = sum
	}
	fmt.Println(fibSer)
}

/* 12. Problem: Program to return a new Array after removing first and last elements of an Array. */
func RFL() {
	nums := []int{11, 12, 13, 14, 15, 16}
	nList := make([]int, 0)
	for i, num := range nums {
		fmt.Println(i, "==", num)
		if i != 0 && i != len(nums)-1 {
			nList = append(nList, num)
		}
	}
	fmt.Println(nList)
}

/* 13. Problem: // Minimum and Maximum element in an Array using sort method. */
func MinMax() {
	nums := []int{7, 11, 6, 55, 98, 45, 16, 96, 46}
	sort.Ints(nums)
	fmt.Println(nums)
	for i, num := range nums {
		if i == 0 {
			fmt.Println("Min Num:", num)
		} else if i == len(nums)-1 {
			fmt.Println("Max Num:", num)
		}
	}
}

/* 14. Problem: Fibonacci Series using HashMap.*/

/* 15. Problem: Sorting of an Array. */

/* 16. Problem: Program to find non-repetitive characters in a string. */

/* 17. Problem: Program to find non-repetitive characters in a string using HashMap. */

/* 18. Problem: Given an array of integers, find the indices of two numbers that add up to a specific target. Assume each input has exactly one solution, and the same element cannot be used twice. */

/* 19. Problem: Program to convert the first and last letter of each word in a string to uppercase. */

/* 20. Problem: Return the maximum profit by choosing a single day to buy and a different day in the future to sell. */

/* 21. Problem: Check if the input is a palindrome. */

/* 22. Problem: Java Program to Expand a String with Character Counts. */

/* 23. Problem: Java Program to Expand a String. */
