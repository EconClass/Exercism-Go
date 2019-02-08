package main

import (
	"fmt"
	// "math"
)

func makeBoard(max int) []uint64 {
	slice := make([]uint64, 64)
	for i := range slice {
		slice[i] = uint64(i+1)
	}

	for i := range slice {
		if slice[i] > 2{
			slice[i] = slice[i-1] << 1
		}
	}

	return slice
}

func Square(index int) uint64 {
	fmt.Println(makeBoard(index))
	return 0
}

func main() {
	Square(64)
}