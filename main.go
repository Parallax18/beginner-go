package main

import "fmt"

var print = fmt.Println
var printf = fmt.Printf

func main() {
	//print(divisibleBy7Not5(2000, 3200))
	//print(factorial(8))
	printf("Mapped squares %v \n", mapNumbersSquared(8))
}
