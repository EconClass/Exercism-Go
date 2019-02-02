package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Valid checks if a number is valid according to the Luhn algorithm
func Valid(input string) bool {
	check := false
	val := strings.TrimSpace(input)

	reg := regexp.MustCompile("\\D")
	val = reg.ReplaceAllString(val, "")
	slice := make([]int, 0)
	for i := range val {
		add, _ := strconv.Atoi(string(val[i]))
		slice = append(slice, add)
	}
	fmt.Println(slice)
	// The first step of the Luhn algorithm is to double every second digit, starting from the right
	// If doubling the number results in a number greater than 9 then subtract 9 from the product
	// Then sum all of the digits; if the sum is evenly divisible by 10, then the number is valid
	for i := len(slice) - 1; i <= 0; i-- {
		fmt.Println(slice[i])
		slice[i] = slice[i] * 2
		fmt.Println(slice[i])
		if slice[i] > 9 {
			slice[i] = slice[i] - 9
		}
	}
	fmt.Println(slice)
	sum := 0
	for i := range slice {
		sum += slice[i]
	}
	fmt.Println(sum)

	if sum != 0 && sum%10 == 0 {
		check = true
	}

	return check
}

func main() {
	fmt.Println(Valid("6103017551"))
}
