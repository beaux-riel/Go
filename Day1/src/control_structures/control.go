package main

import "fmt"

// func main() {
//   fmt.Println(1)
//   fmt.Println(2)
//   fmt.Println(3)
//   fmt.Println(4)
//   fmt.Println(5)
//   fmt.Println(6)
//   fmt.Println(7)
//   fmt.Println(8)
//   fmt.Println(9)
//   fmt.Println(10)
// }

// Versus
// func main() {
// 	fmt.Println(`1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10`)
// }

// I hate vague variables assignments
// func main() {
// 	number := 1
// 	for number <= 10 {
// 		fmt.Println(number)
// 		number = number + 1
// 	}
// }

// Could also be written like this:
// func main() {
// 	for number := 1; number <= 10; number++ {
// 	  fmt.Println(number)
// 	}
//   }

// func main() {
// 	for number := 1; number <= 10; number++ {
// 		if number%2 == 0 {
// 			fmt.Println(number, "even")
// 		} else {
// 			fmt.Println(number, "odd")
// 		}
// 	}
// }

// Switch example
// switch i {
// 	case 0: fmt.Println("Zero")
// 	case 1: fmt.Println("One")
// 	case 2: fmt.Println("Two")
// 	case 3: fmt.Println("Three")
// 	case 4: fmt.Println("Four")
// 	case 5: fmt.Println("Five")
// 	default: fmt.Println("Unknown Number")
// 	}

// Divisible by 3, between 1-100
// func main() {
// 	for number := 1; number <= 100; number++ {
// 		if number%3 == 0 {
// 			fmt.Println(number)
// 		}
// 	}
// }

// Write a program that prints the numbers from 1 to 100.
// But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz".
// For numbers which are multiples of both three and five print "FizzBuzz".

// Ah, the classic FizzBuzz
func main() {
	for number := 1; number <= 100; number++ {
		if number%3 == 0 && number%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if number%3 == 0 {
			fmt.Println("Fizz")
		} else if number%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(number)
		}
	}
}
