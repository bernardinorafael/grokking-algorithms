package main

import "fmt"

func main() {
	for i := range 10 {
		fmt.Printf("Result(%d) = %d\n", i+1, fibonacci(i+1))
	}

}

func fibonacci(n int) int {
	// base case
	if n <= 1 {
		return n
	}
	// recursive case
	return fibonacci(n-1) + fibonacci(n-2)
}
