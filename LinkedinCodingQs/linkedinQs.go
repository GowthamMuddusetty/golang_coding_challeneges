package linkedin

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 1. Problem: Generating 4-digit random number
func GenRandom4Digit() {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 10000
	randNum := rand.Intn(max-min) + min
	if randNum >= 1000 && randNum <= 9999 {
		fmt.Println(randNum)
	}
}

// 2. Problem: Identify the count of a letter in the sentence
func CountLetterOccurence(str string) string {
	runes := []rune(strings.ToLower(str))
	strMap := make(map[string]int)
	for i := 0; i < len(runes); i++ {
		char := string(runes[i])

		//In Go, strMap[char]++ is a shorthand for incrementing the value associated with the key char in the strMap by 1.
		//If the key char doesn't exist in the map, it will be initialized with the zero value for its type, which is 0 for integers.
		// If it already exists, it will be incremented by 1.This is a concise way to update the count of occurrences for each character in the string.
		strMap[char]++
	}

	var result string
	for char, count := range strMap {
		result += fmt.Sprintf("%s:%d ", char, count)
	}

	return result
}
