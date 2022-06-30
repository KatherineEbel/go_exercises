package main

// Grow Given a non-empty array of integers,
// return the result of multiplying the values together in order.
func Grow(arr []int) int {
	result := 1
	for _, val := range arr {
		result *= val
	}
	return result
}
