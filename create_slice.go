package main

import (
	"strconv"
	"strings"
)

// Write a program which accepts a sequence of comma-separated numbers from console and generate an slice out of them. Return the slice.
func createSlice(val string) []int {

	//splits string into slice of strings
	numbers := strings.Split(val, ",")

	length := len(numbers)
	var num = make([]int, length)       //make an empty slice of the length of the array
	for index, value := range numbers { //loop through numbers with range
		num[index], _ = strconv.Atoi(strings.Trim(value, " ")) //update num with integer values after removing whitespace
	}

	return num
}
