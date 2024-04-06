package main

import "math"

func mapNumbersSquared(number int) map[int]int {
	m := make(map[int]int)

	for i := 1; i <= number; i++ {
		m[i] = int(math.Pow(float64(i), 2))
	}

	return m
}
