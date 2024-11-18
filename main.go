package main

import (
	codingQs "codingquestions/codingqs"
)

type Resources struct {
	CodingQs codingQs.GolangCodingQsInterface
}

func (r *Resources) Inject(CodingQs codingQs.GolangCodingQsInterface) {
	r.CodingQs = CodingQs
}

func getServices() *codingQs.CodingQsService {
	return &codingQs.CodingQsService{}
}

func main() {

	tempService := Resources{}
	tempService.Inject(getServices())
	// tempService.CodingQs.GenRandom4Digit()
	// tempService.CodingQs.LinkedListEx()
	// tempService.CodingQs.MutexEx()
	// tempService.CodingQs.SortEx()
	// num := tempService.CodingQs.ReverseNum(1234)
	// fmt.Println(num)
	// tempService.CodingQs.ListOddAndEvenUsingGoroutine()
	// tempService.CodingQs.LongestSubStringFinder()
	// tempService.CodingQs.LongestPalindromicSubstring()
	// tempService.CodingQs.StringReverse("abcd")
	tempService.CodingQs.ZigZagConversion()
}
