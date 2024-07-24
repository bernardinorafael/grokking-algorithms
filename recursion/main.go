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
		fmt.Printf("result = %d\n", fibonacci(i))
	}

	/*
	* result(0) = 0
	* result(1) = 1
	* result(2) = 1
	* result(3) = 2
	* result(4) = 3
	* result(5) = 5
	* result(6) = 8
	* result(7) = 13
	* result(8) = 21
	* result(9) = 34
	 */
}

func fibonacci(n int) int {
	// base case
	if n <= 1 {
		return n
	}
	// recursive case
	return fibonacci(n-1) + fibonacci(n-2)
}
