package enumex

import "fmt"

func Enum() {

	var weekday = Sunday
	fmt.Println(weekday)
	fmt.Println(weekday.String())
	fmt.Println(weekday.EnumIndex())
	fmt.Println("---------------------------------------------------------------------")
}

type Weekday int

const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w Weekday) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[w-1]
}

func (w Weekday) EnumIndex() int {
	return int(w)
}
