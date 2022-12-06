package main // Package name declaration

import "fmt" // Import "fmt" for input/output stuff

var (
	name string = "Beaux"
	address string = "123 Make Believe Street"
)

func main() { // Main function, no surprise

	// name := "Fred" // <- Can only use this style of assignment within the function body
	fmt.Println("Hello " + name + "!")
	fmt.Println("You live at: " + address)
}

// Runs with: beauxwalton$ go run ./name_address.go