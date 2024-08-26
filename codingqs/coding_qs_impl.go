package codingqs

import (
	"cmp"
	"fmt"
	"math"
	"math/rand"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

type CodingQsService struct {
	CodingQsInterface GolangCodingQsInterface
}

/* 1. Problem: Generating 4-digit random number */
func (c *CodingQsService) GenRandom4Digit() {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 10000
	randNum := rand.Intn(max-min) + min
	if randNum >= 1000 && randNum <= 9999 {
		fmt.Println(randNum)
	}
}

/* 2. Problem: Identify the count of a letter in the sentence */
func (c *CodingQsService) CountLetterOccurence(str string) string {
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
func (c *CodingQsService) GenRandom4DigitNoRand() {
	rand.Seed(time.Now().UnixNano())

	// generate a random 8-digit ID
	id := ""
	for i := 0; i < 8; i++ {
		id += strconv.Itoa(rand.Intn(10))
	}
	fmt.Println("---------------------------------------------------------------------")
}

/* 5. Problem: Finding number in an Array and ArrayList */
func (c *CodingQsService) Find_Num() {
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
func (c *CodingQsService) Highest_Num() {
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
func (c *CodingQsService) Char_Count() {
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
func (c *CodingQsService) Extract_Char() {
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
func (c *CodingQsService) RemDup() {
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
func (c *CodingQsService) FirstLet() {
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
func (c *CodingQsService) FibSeries() {
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
func (c *CodingQsService) RFL() {
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
func (c *CodingQsService) MinMax() {
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
func (c *CodingQsService) FibMap() {
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
func (c *CodingQsService) ArrSort() {
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
func (c *CodingQsService) NonRepetetiveChars() {
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
func (c *CodingQsService) NonRepetetiveCharsUsingMap() {
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
func (c *CodingQsService) IndicesOfSumOf2Nums() {
	nums := []int{7, 11, 6, 55, 98, 45, 16, 96, 46, 7, 0}
	nonRepNums := make([]int, len(nums))
	targetSum := 22

	for i, num := range nums {
		if num != 0 {
			for j := 0; j < len(nonRepNums); j++ {
				if num == nonRepNums[j] {
					break
				} else {
					nonRepNums[i] = num
				}
			}
		}
	}

	filteredNums := make([]int, 0)
	for _, num := range nonRepNums {
		if num != 0 {
			filteredNums = append(filteredNums, num)
		}
	}

	slices.Sort(filteredNums)
	indices := make([]int, 0)
	for _, num := range filteredNums {
		diff := targetSum - num
		for i := 0; i < len(filteredNums); i++ {
			if diff == filteredNums[i] && num != diff {
				indices = append(indices, i)
			}
		}
	}

	slices.Sort(indices)
	fmt.Println(indices)
	fmt.Println("---------------------------------------------------------------------")
}

/* 19. Problem: Program to convert the first and last letter of each word in a string to uppercase. */
func (c *CodingQsService) FirstLastLetterUpperCase() {
	str := "success is a not final, failure is not fatal: it is the courage to continue that counts."

	words := strings.Split(str, " ")
	for i, word := range words {
		runes := []rune(word)
		firstRune := unicode.ToUpper(runes[0])
		lastRune := unicode.ToUpper(runes[len(word)-1])
		runes[0] = firstRune
		wordLen := len(runes)
		if wordLen == 1 {
			fmt.Println(string(runes))
			words[i] = string(unicode.ToUpper(runes[0]))
		}
		if !unicode.IsLetter(runes[len(word)-1]) {
			lastRune = unicode.ToUpper(runes[len(word)-2])
			runes[len(word)-2] = lastRune
			alteredWords := string(runes)
			words[i] = alteredWords
		} else {
			runes[len(word)-1] = lastRune
			alteredWords := string(runes)
			words[i] = alteredWords
		}
	}
	fmt.Println(strings.Join(words, " "))
}

/* 20. Problem: Return the maximum profit by choosing a single day to buy and a different day in the future to sell. */
func (c *CodingQsService) CalculateMaxProfit() {
	prices := []int{7, 5, 2, 3, 6, 4, 1}

	profitCalFunc := func([]int) int {
		if len(prices) <= 1 {
			return 0
		}

		minPrice := prices[0]
		maxProfit := 0

		for i := 1; i < len(prices); i++ {
			if prices[i] < minPrice {
				minPrice = prices[i]
			} else {
				profit := prices[i] - minPrice
				if profit > maxProfit {
					maxProfit = profit
				}
			}
		}
		return maxProfit
	}

	profit := profitCalFunc(prices)
	if profit > 0 {
		fmt.Println("Profit Earned:", profit)
	} else {
		fmt.Println("No Profit Earned")
	}
}

/* 21. Problem: Check if the input is a palindrome. */
func (c *CodingQsService) CheckPalindrome() {
	fmt.Println("Enter input :")
	var input string
	fmt.Scan(&input)
	isInt := func(s string) bool {
		for _, c := range s {
			if !unicode.IsDigit(c) {
				return false
			}
		}
		return true
	}

	reverseFunc := func(strs []interface{}) chan interface{} {
		ret := make(chan interface{})
		go func() {
			for i, _ := range strs {
				ret <- strs[len(strs)-1-i]
			}
			close(ret)
		}()
		return ret
	}

	checkIntPalindrome := func(num int) bool {
		isTrue := false
		if num <= 0 {
			fmt.Println("Enter valid number greater than 0: ", num)
		} else {
			expectedNum := num
			actualNum := 0

			for num != 0 {
				n := num % 10
				actualNum = actualNum*10 + n
				num = num / 10
			}

			if expectedNum == actualNum {
				isTrue = true
			}
		}
		return isTrue
	}

	checkStringPalindrome := func(str string) bool {
		actualStr := str
		expectedStr := ""
		isTrue := false
		var arrStr []interface{}
		for _, c := range input {
			arrStr = append(arrStr, string(c))
		}
		ch := reverseFunc(arrStr)
		revStr := make([]string, 0)
		for val := range ch {
			revStr = append(revStr, val.(string))
		}
		expectedStr = strings.Join(revStr, "")

		if actualStr == expectedStr {
			isTrue = true
		}
		return isTrue
	}

	if isInt(input) {
		num, _ := strconv.Atoi(input)
		value := checkIntPalindrome(num)
		if value {
			fmt.Println("Num ", num, " is palindrome")
		} else {
			fmt.Println("Num ", num, " is not palindrome")
		}
	} else {
		val := checkStringPalindrome(input)
		if val {
			fmt.Println("String ", input, " is palindrome")
		} else {
			fmt.Println("String ", input, " is not palindrome")
		}
	}
}

/* 22. Problem:  Program to Expand a String with Character Counts. */
func (c *CodingQsService) StringExpander() {
	inputString := "a3.2b2@c3!2d?"
	expandedString := ""
	for i, s := range inputString {
		if !unicode.IsDigit(s) {
			expandedString = expandedString + string(s)
			if i+1 < len(inputString) && unicode.IsDigit(rune(inputString[i+1])) {
				num, _ := strconv.Atoi(string(inputString[i+1]))
				for num > 1 {
					expandedString = expandedString + string(s)
					num--
				}
			}

		}
	}
	fmt.Println(expandedString)
}



/* 23. Problem:  Program to Expand a String. */
func (c *CodingQsService) StringExpander2() {
	inputString := "a3.2B2@c3!2d?"
	expandedString := ""
	for i, s := range inputString {
		if !unicode.IsDigit(s) {
			expandedString = expandedString + string(s)
			if i+1 < len(inputString) && unicode.IsDigit(rune(inputString[i+1])) {
				num, _ := strconv.Atoi(string(inputString[i+1]))
				for num > 1 {
					expandedString = expandedString + string(s)
					num--
				}
			}

		}
	}
	fmt.Println(expandedString)
}

func (cqs *CodingQsService) MaxLenOfSubString() {

}

/* LINKED LIST DATA STRUCTURE */
func (c *CodingQsService) LinkedListEx() {
	type LinkedListTest struct {
		Num  int
		Next *LinkedListTest
	}

	head := &LinkedListTest{Num: 1}
	current := head
	for i := 2; i <= 10; i++ {
		current.Next = &LinkedListTest{Num: i}
		current = current.Next
	}

	current = head
	for current != nil {
		fmt.Println(current.Num)
		current = current.Next
	}
}

/*  MUTEX_EX */
func (c *CodingQsService) MutexEx() {

	var counter int
	var counterMutex sync.Mutex
	ch := make(chan string, 1)
	increment := func() {
		counterMutex.Lock()
		defer counterMutex.Unlock()
		counter++
		if counter == 1000 {
			ch <- "done"
		}
	}

	// var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		// wg.Add(1)
		go func() {
			// defer wg.Done()
			increment()
		}()

	}

	// wg.Wait()
	<-ch
	fmt.Println("Final counter value:", counter)
}

func (c *CodingQsService) SortEx() {
	fruits := []string{"kiwi", "peach", "watermelon"}
	sortByLen := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, sortByLen)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "Gowtham", age: 23},
		{name: "Jagadeesh", age: 32},
	}
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)
}

func (c *CodingQsService) ReverseNum(x int) int {
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
