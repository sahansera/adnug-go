package main

import "fmt"

type Person struct {
	name string
	age  int
}

// Q: How can we add behaviours to this?

func main() {
	person := Person{}
	fmt.Println(person)
}
