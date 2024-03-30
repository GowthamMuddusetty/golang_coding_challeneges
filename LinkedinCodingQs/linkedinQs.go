package linkedin

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/*
1. Problem: Generating 4-digit random number
*/
func GenRandom4Digit() {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 10000
	randNum := rand.Intn(max-min) + min
	if randNum >= 1000 && randNum <= 9999 {
		fmt.Println(randNum)
	}
}

/*
2. Problem: Identify the count of a letter in the sentence
*/
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
