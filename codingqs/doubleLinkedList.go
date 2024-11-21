package codingqs

import "fmt"

type Node[T any] struct {
	Prev  *Node[T]
	Value T
	Next  *Node[T]
}

type DoubleLinkedList[T any] struct {
	Head, Tail *Node[T]
}

func (dll *DoubleLinkedList[T]) Push(v T) {
	newNode := &Node[T]{Value: v}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = dll.Head
		return
	}

	dll.Tail.Next = newNode
	newNode.Prev = dll.Tail
	dll.Tail = newNode
}

func (dll *DoubleLinkedList[T]) StoreAllValues() []T {

	var val []T
	for e := dll.Head; e != nil; e = e.Next {
		fmt.Println("Node: ", e)
		val = append(val, e.Value)
	}
	return val
}

func PrintDLLValues() {
	obj := &DoubleLinkedList[int]{}
	obj.Push(10)
	obj.Push(20)
	obj.Push(30)
	obj.Push(40)
	obj.Push(50)
	fmt.Println(obj.StoreAllValues())
}
