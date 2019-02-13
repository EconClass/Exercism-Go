package grains

import (
	"errors"
)

func makeBoard(max int) []uint64 {
	slice := make([]uint64, 64)
	for i := range slice {
		slice[i] = uint64(i+1)
		if slice[i] > 2 {
			slice[i] = slice[i-1] << 1
		}
	}
	return slice
}

func Square(index int) (uint64, error) {
	if index < 1 || index > 64 {
		return 0, errors.New("Pick a number from 1 to 64")
	}
	if index <= 2 {
		return uint64(index), nil
	}
	return 1 << uint64(index-1), nil
}

func Total() uint64{
	var board []uint64
	board = makeBoard(64)
	var sum uint64 = 0
	for i := range board {
		sum += board[i]
	}
	return sum
}