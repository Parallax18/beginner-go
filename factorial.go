package main

import "strconv"

func factorial(number int) string {
	var f int
	for i := number; i >= 0; i-- {
		f *= i
	}

	return strconv.FormatInt(int64(f), 10)
}
