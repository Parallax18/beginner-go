package main

//Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
//
//Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
//I can be placed before V (5) and X (10) to make 4 and 9.
//X can be placed before L (50) and C (100) to make 40 and 90.
//C can be placed before D (500) and M (1000) to make 400 and 900.
//Given a roman numeral, convert it to an integer.

var romanNumerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {

	var valueSlice []int
	for _, character := range s {
		valueSlice = append(valueSlice, romanNumerals[string(character)])
	}
	print(valueSlice)
	var total int
	toZeroIndexes := make(map[int]int)
	for index, val := range valueSlice {
		if index > 0 && val > valueSlice[index-1] {

			valueSlice[index-1] = val - valueSlice[index-1]
			toZeroIndexes[index] = index

		}

	}

	for index, _ := range valueSlice {
		if _, isFound := toZeroIndexes[index]; isFound {
			valueSlice[toZeroIndexes[index]] = 0
			print(toZeroIndexes[index], isFound)
		}
		print(valueSlice)
	}

	for _, num := range valueSlice {
		total += num
	}

	return total
}

//Optimal Solution
//func romanToInt(s string) int {
//	var total, prevValue int
//	for _, character := range s {
//		value := romanNumerals[string(character)]
//		total += value
//		if value > prevValue {
//			total -= 2 * prevValue // Subtract twice the previous value
//		}
//		prevValue = value
//	}
//	return total
//}
