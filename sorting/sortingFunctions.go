package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
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

	// map1 := linkedin.CountLetterOccurence("Hello")
	// fmt.Println(map1)

}
