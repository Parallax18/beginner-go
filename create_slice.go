package main

import "strings"

func createSlice(val string) []string {
	return strings.Split(val, ",")
}
