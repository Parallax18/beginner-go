package main

import (
	"strconv"
	"strings"
)

//	Write a program which will find all such numbers which are divisible by 7 but are not a multiple of 5, between 2000 and 3200 (both included). The numbers obtained should be printed in a comma-separated sequence on a single line.

func divisibleBy7Not5(lower, upper int) string {

	var values []string
	for i := lower; i <= upper; i++ {
		if i%7 == 0 && i%5 != 0 {
			values = append(values, strconv.Itoa(i))

		}
	}
	return strings.Join(values, ",")

}
