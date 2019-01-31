package diffsquares

// NatNum crates a sequence of natural numbers
func NatNum(bound int) []int {
	naturals := make([]int, bound)
	for i := range naturals {
		naturals[i] = i + 1
	}
	return naturals
}

// SquareOfSum adds all natural numbers in a sequence then squares the sum
func SquareOfSum(max int) int {
	sequence := NatNum(max)
	prod := 0
	for i := range sequence {
		prod += sequence[i]
	}
	prod = prod * prod
	return prod
}

// SumOfSquares squares all the natural numbers in a sequence then sums the squares
func SumOfSquares(max int) int {
	sequence := NatNum(max)
	sum := 0
	for i := range sequence {
		sum += sequence[i] * sequence[i]
	}
	return sum
}

// Difference finds the difference between the the results of SquareOfSum and SumOfSquares functions
func Difference(max int) int {
	return SquareOfSum(max) - SumOfSquares(max)
}
