package main

import (
	linkedin "codingQuestions/linkedinCodingQs"
)

func getServices() *linkedin.CodingQsService {
	return &linkedin.CodingQsService{}
}

func main() {

	// MUTEX -- execute below func
	// mutex.MutexEx()

	// LINKED LIST STRUCTURE -- execute below func
	// linkedlist.LinkedList()

	// ENUM -- execute below func
	// enumex.Enum()

	tempService := linkedin.CodingQsService{}
	tempService.Inject(getServices())

	// tempService.Find_Num()
	tempService.Highest_Num()
	// tempService.Char_Count()
	// tempService.Extract_Char()
	// tempService.RemDup()
	// tempService.FirstLet()
	// tempService.FibSeries()
	// tempService.RFL()
	// tempService.MinMax()
	// tempService.FibMap()
	// tempService.ArrSort()
	// tempService.NonRepetetiveChars()
	// tempService.NonRepetetiveCharsUsingMap()
	// tempService.IndicesOfSumOf2Nums()
	// tempService.FirstLastLetterUpperCase()
	// tempService.CalculateMaxProfit()
	// tempService.CheckPalindrome()
}
