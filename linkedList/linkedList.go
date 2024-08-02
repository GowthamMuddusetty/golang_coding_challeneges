package linkedlist

import "fmt"

type LinkedListTest struct {
	Num  int
	Next *LinkedListTest
}

func LinkedList() {

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
