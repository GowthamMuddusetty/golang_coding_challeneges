package main

import (
	codingQs "codingquestions/codingqs"
)

type Resources struct {
	CodingQs   codingQs.GolangCodingQsInterface
	LeetCodeQs codingQs.LeetCodeQuestionsInterface
}

func (r *Resources) Inject(CodingQs codingQs.GolangCodingQsInterface, LeetCodeQs codingQs.LeetCodeQuestionsInterface) {
	r.CodingQs = CodingQs
	r.LeetCodeQs = LeetCodeQs
}

func getServices() (*codingQs.CodingQsService, *codingQs.LeetCodeQsService) {
	return &codingQs.CodingQsService{}, &codingQs.LeetCodeQsService{}
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
	tempService.LeetCodeQs.AddTwoNumbers()
}
