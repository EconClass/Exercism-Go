package main

import (
	"fmt"
	"math"
)

func Square(upper int) uint64 {
	slice := make([]float64, 64)
	for i := range slice {
		slice[i] = math.Pow(float64(2), float64(i))
	}
	fmt.Println(slice)
	return 0
}

func main() {
	Square(64)
}