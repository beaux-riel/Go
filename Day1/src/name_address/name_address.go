package main // Package name declaration

import "fmt" // Import "fmt" for input/output stuff

var name string = "Beaux"
var address string = "123 Make Believe Street"

func main() { // Main function, no surprise
	fmt.Println("Hello " + name + "!")
	fmt.Println("You live at: " + address)
}

// Runs with: beauxwalton$ go run ./name_address.go