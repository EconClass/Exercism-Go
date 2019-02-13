package grains

import (
	"errors"
)

// MakeBoard creates an slice of uint64 where each successive number is double the previous
func MakeBoard(max int) []uint64 {
	slice := make([]uint64, max)
	for i := range slice {
		slice[i] = 1 << uint64(i)
	}
	return slice
}

// Square returns a uint64 equal to the number of grains are in the chess board square
func Square(index int) (uint64, error) {
	if index < 1 || index > 64 {
		return 0, errors.New("Pick a number from 1 to 64")
	}
	return 1 << uint64(index-1), nil
}

// Total creates a slice containing 64 values using MakeBoard, then sums all values in the slice
func Total() uint64 {
	var board []uint64
	board = MakeBoard(64)
	var sum uint64 = 0
	for _, grains := range board {
		sum += grains
	}
	return sum
}