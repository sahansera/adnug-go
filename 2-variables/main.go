package main

import "fmt"

func main() {
	// Declaring variables
	var myInt int = 1
	var myFloat float32 = 2.0
	var myString string = "Hello"
	fmt.Println(myInt, myFloat, myString)

	// Declaring multiple variables
	var (
		bookId        int    = 3
		title, author string = "Programming Go", "Unknown"
		published     bool   = false
	)
	fmt.Println(bookId, title, author, published)

	// Type inferencing
	var other = "other"
	fmt.Println(other)

	// Shorter version
	anotherBookId := 4
	anotherTitle, anotherAuthor := "Another", "Unknown"
	fmt.Println(anotherBookId, anotherTitle, anotherAuthor)

	// Constants
	// Declaration of constants is similar to that of variables
	const fixedStr string = "Not changing"
	const fixedInt = 1
	const (
		fixedBool   = false
		fixedDouble = 10.0
	)
}
