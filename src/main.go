package main

import "fmt"

// Sum calculates the sum of two integers
func Sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Sum:", Sum(1, 2))
}
