package codingqs

type GolangCodingQsInterface interface {

	/* 1. Problem: Generating 4-digit random number */
	GenRandom4Digit()

	/* 2. Problem: Identify the count of a letter in the sentence */
	CountLetterOccurence(str string) string

	/* 4. Generate random without using math.rand function */
	GenRandom4DigitNoRand()

	/* 5. Problem: Finding number in an Array and ArrayList */
	Find_Num()

	/* 6. Problem: Find out highest numbers in an Array whose adjacent right numbers are smaller */
	Highest_Num()

	/*7. Problem: Find a count of each character in a given String.*/
	Char_Count()

	/* 8. Problem: Extract last 4 characters of any String.*/
	Extract_Char()

	/* 9. Problem: Program to remove duplicates from a given array. */
	RemDup()

	/* 10. Problem: First letter of each word present in String. */
	FirstLet()

	/* 11. Problem: Fibonacci Series using Arrays. */
	FibSeries()

	/* 12. Problem: Program to return a new Array after removing first and last elements of an Array. */
	RFL()

	/* 13. Problem: // Minimum and Maximum element in an Array using sort method. */
	MinMax()

	/* 14. Problem: Fibonacci Series using HashMap.*/
	FibMap()

	/* 15. Problem: Sorting of an Array. */
	ArrSort()

	/* 16. Problem: Program to find non-repetitive characters in a string. */
	NonRepetetiveChars()

	/* 17. Problem: Program to find non-repetitive characters in a string using HashMap. */
	NonRepetetiveCharsUsingMap()

	/* 18. Problem: Given an array of integers, find the indices of two numbers that add up to a specific target. Assume each input has exactly one solution, and the same element cannot be used twice. */
	IndicesOfSumOf2Nums()

	/* 19. Problem: Program to convert the first and last letter of each word in a string to uppercase. */
	FirstLastLetterUpperCase()

	/* 20. Problem: Return the maximum profit by choosing a single day to buy and a different day in the future to sell. */
	CalculateMaxProfit()

	/* 21. Problem: Check if the input is a palindrome. */
	CheckPalindrome()

	/* 22. Problem:  Program to Expand a String with Character Counts. */
	StringExpander()

	/* 23. Problem:  Program to Expand a String. */
	StringExpander2()

	/* 24. Problem: Program to Find the Longest Contiguous Alphanumeric Substring. */
	LongestSubStringFinder()

	/* 26. Problem: Program to Find the Longest Palindromic Substring.*/
	LongestPalindromicSubstring()

	/* Given a signed 32-bit integer x, return x with its digits reversed.
	If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0." */
	ReverseNum(x int) int

	/* divide an array in 2 diff goroutines odd and even */
	ListOddAndEvenUsingGoroutine()

	/* LINKEDLIST */
	LinkedListEx()

	// MUTEX_EX
	MutexEx()

	// SORT_EX
	SortEx()
}
