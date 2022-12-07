package main

import "fmt"

// func main() {
// 	var my_array [5]int
// 	my_array[0] = 7
// 	my_array[4] = 100
// 	fmt.Println(my_array)
// }

func main() {
	var scores [5]float64
	scores[0] = 98
	scores[1] = 93
	scores[2] = 77
	scores[3] = 82
	scores[4] = 83

	var total float64 = 0
	for number := 0; number < 5; number++ {
		total += scores[number]
	}
	fmt.Println("The average score is:", total/5) // Notice the comma between the string and the float?
}
