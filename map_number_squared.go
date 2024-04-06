package main

import "math"

//With a given integral number n, write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included), and then the program should print the map with representation of the value

func mapNumbersSquared(number int) map[int]int {
	m := make(map[int]int)

	for i := 1; i <= number; i++ {
		m[i] = int(math.Pow(float64(i), 2))
	}

	return m
}
