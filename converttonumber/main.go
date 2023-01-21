package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	maplist := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	splitinput := strings.Split(input, "")

	var total int
	for _, rome := range splitinput {
		number, exist := maplist[rome]

		if !exist {
			panic("The string is not exist")
		}

		total += number
	}

	fmt.Println(total)
}
