package clock

import (
	"fmt"
)

// Implement a clock that handles cs without dates.
// You should be able to add and subtract minutes to it.
// Two clocks that represent the same c should be equal to each other.

// Clock is used to keep track of hours and minutes
type Clock struct {
	hours   int
	minutes int
}

// New Returns a new Clock type given the hours and minutes field
func New(hours, minutes int) Clock {
	return Clock{hours, minutes}
}

func (c Clock) String() string {
	result := "%v:%v"
	if c.hours > 9 && c.minutes > 9 {
		return fmt.Sprintf(result, c.hours, c.minutes)
	}
	return fmt.Sprintf(result, c.hours, c.minutes)
}

// Add adds minutes to an existing Clock type
func (c Clock) Add(add int) Clock {
	c.minutes = c.minutes + add
	fmt.Println(c.minutes)
	addHours := 0
	for c.minutes > 60 {
		c.minutes = c.minutes - 60
		addHours++
	}
	c.hours = c.hours + addHours
	return c
}

// Subtract subtracts minutes from an existing Clock type
func (c Clock) Subtract(subtract int) Clock {
	c.minutes = c.minutes - subtract
	removeHours := 0
	for c.minutes < 0 {
		c.minutes = c.minutes + 60
		removeHours--
	}
	c.hours = c.hours + removeHours
	return c
}

// CompareClocks checks if the c indidcated by two Clock types are the same
func CompareClocks(first, second Clock) bool {
	if first.hours != second.hours || first.minutes != second.minutes {
		return false
	}
	return true
}

// func main() {
// 	clock := New(8, 0)
// 	// fmt.Println(clock.String())
// }
