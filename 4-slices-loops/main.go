package main

import "fmt"

func main() {
	// Arrays are fixed length
	var array = [3]int{0, 2, 3}
	array[0] = 1

	// Slices are dynamic
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println(array, numbers)

	// Loops
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	// Loop through the slice
	for i, n := range numbers {
		fmt.Println(i, n)
	}
}
