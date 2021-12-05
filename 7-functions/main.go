package main

import "fmt"

func main() {
	getGreeting()
	sum, avg := sumAndAvg(1, 1)
	fmt.Printf("sum:%d, avg:%d\n", sum, avg)
}

// A simple example with a single return value
func getGreeting() string {
	return "Hello"
}

// An example with multiple return values
// This is also commonly used when a function
// can also return an error
func sumAndAvg(a, b int) (int, int) {
	return (a + b), (a + b) / 2
}
