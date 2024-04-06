package main

import "strconv"

// Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.
func factorial(number int) string {
	var f int = 1 // default value of int is Zero, initialize as 1 to avoid zero multiplication.
	for i := number; i > 1; i-- {
		f *= i
	}

	return strconv.FormatInt(int64(f), 10)
}
