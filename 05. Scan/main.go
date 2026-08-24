package main

import "fmt"

func main() {
	var choice int;
	fmt.Printf("Enter your number: ")
	fmt.Scan(&choice)

	fmt.Printf("The number is: %d", choice)
}