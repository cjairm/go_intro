// Package algorithm contains different algorithms
package algorithm

// FactorialNumber create a factorial of a number
// Create a function called "FactorialNumber"
func FactorialNumber(n int) int {
	if n == 1 {
		return 1
	} else if n <= 0 {
		return 0
	}
	return n * FactorialNumber(n-1)
}
