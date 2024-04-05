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

    // Check for empty input
    if strings.TrimSpace(userInput) == "" {
        fmt.Println("Error: Input cannot be empty.")
        return
    }

    fmt.Println("You entered:", userInput)
}
