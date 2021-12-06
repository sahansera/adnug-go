package main

import "fmt"

func main() {
	// Memory Addresses
	// This can be accessed with the ampersand (&) operator
	var x int = 5
	fmt.Printf("Memory Address of variable x: %x\n", &x)
	fmt.Printf("Value of variable x: %d\n", x)

	// Pointers
	// A pointer is a variable whose address is the direct memory address of another variable
	var ipointer *int = &x
	fmt.Printf("Memory Address of variable x: %x\n", ipointer)
}
