package main

import (
	"fmt"
	"strings"
)

func main() {
	var userInput string

	// Taking user input
	fmt.Print("Enter your name: ")
	fmt.Scanln(&userInput)

	// Displaying user input
	fmt.Println("Hello,", userInput)
}
