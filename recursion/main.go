package main

import "fmt"

/*
* Fibonacci
* F(0) = 0
* F(1) = 1
* F(n) = F(n-1) + F(n-2)
 */

func main() {
	for i := range 10 {
		fmt.Printf("Result(%d) = %d\n", i, fibonacci(i))
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
