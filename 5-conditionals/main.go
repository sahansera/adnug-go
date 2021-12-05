package main

import "fmt"

func main() {

	// Conditionals syntax is similar to C#
	var val = 10
	if val < 10 {
		fmt.Println("less than 10")
	} else if val >= 5 && val <= 10 {
		fmt.Println("between 5 and 10")
	} else {
		fmt.Println("greater than 10")
	}

	// Switch expressions
	var lang = "go"
	switch lang {
	case "go":
		fmt.Println("go run")
	case "c#":
		fmt.Println("dotnet run")
	default:
		fmt.Println("Invalid language")
	}
}
