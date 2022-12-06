// Example code from: https://www.golang-book.com/books/intro/4

package main

import "fmt"

// func main() {
//   var x string
//   x = "first"
//   fmt.Println(x)
//   x = "second"
//   fmt.Println(x)
// }

func main() {
	var x string
	x = "first "
	fmt.Println(x)
	x += "second "
	fmt.Println(x)
	x += "third "
	fmt.Println(x)
	x += x
	fmt.Println(x)
}