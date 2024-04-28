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
	fmt.Println("---------------------------------------------------------------------")
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
	fmt.Println("---------------------------------------------------------------------")
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
	fmt.Println("---------------------------------------------------------------------")
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
	fmt.Println("---------------------------------------------------------------------")
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
	fmt.Println("---------------------------------------------------------------------")
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
	fmt.Println("---------------------------------------------------------------------")
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
	fmt.Println("---------------------------------------------------------------------")
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
	fmt.Println("---------------------------------------------------------------------")
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
	fmt.Println("---------------------------------------------------------------------")
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
	fmt.Println("---------------------------------------------------------------------")
}

/* 14. Problem: Fibonacci Series using HashMap.*/
func FibMap() {
	num := 10
	hashmap := make(map[int]int, num)
	hashmap[0] = 0
	hashmap[1] = 1

	for i := 2; i < num; i++ {
		hashmap[i] = hashmap[i-1] + hashmap[i-2]
	}

	//keys := make([]int, 0)
	keys := make([]int, 0, len(hashmap))
	for key := range hashmap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for key := range keys {
		fmt.Println(key, "--", hashmap[key])
	}
	fmt.Println("---------------------------------------------------------------------")
}

/* 15. Problem: Sorting of an Array. */
func ArrSort() {
	arr := []int{7, 11, 65, 98, 45, 16, 96, 6}
	sortedArr := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			newNum := arr[j]
			if arr[i] > newNum {
				count = count + 1
			}
		}
		sortedArr[count] = arr[i]
	}
	fmt.Println("Sorted using sort method...", sortedArr)

	// 2nd way
	// sort.Ints(arr)
	// fmt.Println("Sorted using sort method...", arr)
	fmt.Println("---------------------------------------------------------------------")
}

/* 16. Problem: Program to find non-repetitive characters in a string. */
func NonRepetetiveChars() {
	str := strings.ToLower("Sunset glow on the horizon")
	runes := []rune(str)
	chars := make([]string, 0)

	for i := 0; i < len(runes); i++ {
		char := runes[i]
		count := 0

		for j := 0; j < len(runes); j++ {
			if char == runes[j] {
				count++
			}
		}

		if count == 1 {
			chars = append(chars, string(char))
		}
	}

	fmt.Println(chars)
	fmt.Println("---------------------------------------------------------------------")
}

/* 17. Problem: Program to find non-repetitive characters in a string using HashMap. */
func NonRepetetiveCharsUsingMap() {
	str := strings.ToLower("Sunset glow on the horizon")
	noRepcharmap := make(map[string]int)

	for _, char := range str {
		noRepcharmap[string(char)]++
	}

	var keysWithValOne []string

	for key, val := range noRepcharmap {
		if val == 1 {
			keysWithValOne = append(keysWithValOne, key)
		}
	}

	fmt.Println(strings.Join(keysWithValOne, ","))
	fmt.Println("---------------------------------------------------------------------")
}

/* 18. Problem: Given an array of integers, find the indices of two numbers that add up to a specific target. Assume each input has exactly one solution, and the same element cannot be used twice. */
func IndicesOfSumOf2Nums() {
	nums := []int{7, 11, 6, 55, 98, 45, 16, 96, 46, 7}
	nonRepNums := make([]int, len(nums))
	targetSum := 22

	for i, num := range nums {
		for j := 0; j < len(nonRepNums); j++ {
			if num == nonRepNums[j] {
				break
			} else {
				nonRepNums[i] = num
			}
		}
	}

	filteredNums := make([]int, 0)
	for _, num := range nonRepNums {
		if num != 0 {
			filteredNums = append(filteredNums, num)
		}
	}

	indices := make([]int, 0)
	for _, num := range filteredNums {
		diff := targetSum - num
		for i := 0; i < len(filteredNums); i++ {
			if diff == filteredNums[i] {
				indices = append(indices, i)
			}
		}
	}

	fmt.Println("---------------------------------------------------------------------")
}

/* 19. Problem: Program to convert the first and last letter of each word in a string to uppercase. */

/* 20. Problem: Return the maximum profit by choosing a single day to buy and a different day in the future to sell. */

/* 21. Problem: Check if the input is a palindrome. */

/* 22. Problem: Java Program to Expand a String with Character Counts. */

/* 23. Problem: Java Program to Expand a String. */
