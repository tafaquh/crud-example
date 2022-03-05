package main

import (
	"fmt"
	"os"
	"strconv"
)

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func main() {
	// Read input
	n := os.Getenv("NUMBER")

	number, _ := strconv.Atoi(n)

	// Run Fibo
	result := FibonacciRecursion(number)

	fmt.Printf("\nFibonacchi of %d is %d\n", number, result)

}
