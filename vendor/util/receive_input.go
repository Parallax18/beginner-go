package util

import (
	"fmt"
	"log"
)

func ReceiveInput(message string) string {
	var input string

	print(message)

	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatalln("Please provide an input")
	}

	return input
}
