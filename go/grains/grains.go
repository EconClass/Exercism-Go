package main

import (
	"fmt"
	"math"
)

func makeBoard(max int) []uint64 {
	slice := make([]uint64, max)
	for i := range slice {
		slice[i] = uint64(math.Pow(float64(2), float64(i)))
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