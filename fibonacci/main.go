package main

import "fmt"

func fibonacci(number int) []int {
	var fibonacci []int
	if number <= 2 {
		return nil
	}

	for i := 1; i <= number; i++ {
		if i == 1 || i == 2 {
			fibonacci = append(fibonacci, 1)
			continue
		}
		fibonacci = append(fibonacci, fibonacci[i-2]+fibonacci[i-3])
	}

	return fibonacci
}

func main() {
	var number int
	fmt.Print("What's the Fibonacci sequence you want? ")
	fmt.Scan(&number)

	fmt.Println("The Fibonacci sequence is:", fibonacci(number))
}
