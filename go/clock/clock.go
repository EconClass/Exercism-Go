package clock

import (
	"fmt"
)

// Clock is used to keep track of hours and minutes
type Clock struct {
	hours   int
	minutes int
}

// New Returns a new Clock type given the hours and minutes field
func New(hours, minutes int) Clock {
	for minutes >= 60 {
		hours++
		minutes = minutes - 60
	}
	for minutes < 0 {
		hours--
		minutes = minutes + 60
	}
	for hours > 23 {
		hours = hours - 24
	}
	for hours < 0 {
		hours = hours + 24
	}

	return Clock{hours, minutes}
}

// String is a method of the Clock type that returns a string of the time in the Clock
func (c Clock) String() string {
	result := ""
	switch {
	case c.hours > 9 && c.minutes > 9:
		result = "%v:%v"
	case c.hours > 9 && c.minutes < 10:
		result = "%v:0%v"
	case c.hours < 10 && c.minutes > 9:
		result = "0%v:%v"
	default:
		result = "0%v:0%v"
	}
	return fmt.Sprintf(result, c.hours, c.minutes)
}

// Add is a method of the Clock type that adds minutes to an existing Clock type
func (c Clock) Add(add int) Clock {
	c.minutes = c.minutes + add
	// fmt.Println(c.minutes)
	addHours := 0
	for c.minutes >= 60 {
		c.minutes = c.minutes - 60
		addHours++
	}
	// fmt.Println(c.minutes)
	c.hours = c.hours + addHours
	for c.hours > 23 {
		c.hours = c.hours - 24
	}
	return c
}

// Subtract is a method of the Clock type that subtracts minutes from an existing Clock type
func (c Clock) Subtract(subtract int) Clock {
	c.minutes = c.minutes - subtract
	removeHours := 0
	for c.minutes < 0 {
		c.minutes = c.minutes + 60
		removeHours--
	}
	c.hours = c.hours + removeHours
	for c.hours < 0 {
		c.hours = c.hours + 24
	}
	return c
}

// CompareClocks checks if the Clocks passed into it are the same
func CompareClocks(first, second Clock) bool {
	if first.hours != second.hours || first.minutes != second.minutes {
		return false
	}
	return true
}
