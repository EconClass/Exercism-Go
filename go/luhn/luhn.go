package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

// Valid checks if a number is valid according to the Luhn algorithm
func Valid(input string) bool {
	check := false
	val := strings.Replace(input, " ", "", -1)

	if val == "0" {
		return false
	}

	slice := make([]int, 0)
	for i := range val {
		add, err := strconv.Atoi(string(val[i]))
		if err != nil {
			fmt.Println(err)
			return false
		}
		slice = append(slice, add)
	}

	for i := len(slice) - 2; i >= 0; i = i - 2 {
		slice[i] = slice[i] * 2

		if slice[i] > 9 {
			slice[i] = slice[i] - 9
		}
	}

	sum := 0
	for i := range slice {
		sum += slice[i]
	}

	if sum%10 == 0 {
		check = true
	}

	return check
}
