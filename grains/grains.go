package main

import (
	"fmt"
	"math"
)

func Square(upper int) float64 {
	valMap := make(map[int]uint64)
	for i := range valMap {
		valMap[i] = uint64(math.Pow(float64(2), float64(i)))
	}
	fmt.Println(valMap)
	return 0
}

func main() {
	Square(64)
}