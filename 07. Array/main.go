package main

import "fmt"

func main() {
	var numbers [6]int
	fmt.Println(numbers)  // [0 0 0 0 0 0]

	numbers[1] = 43;
	numbers[3] = 25;
	numbers[1] = 61;
	fmt.Println(numbers) 	// [0 61 0 25 0 0]

	fmt.Println("Length of the array:", len(numbers)) // Length of the array: 6

	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}

	//* ------------------------------------------ *//
	values := [5]int {6, 9, 3, 5, 7}
	fmt.Print("Values:", values)
}