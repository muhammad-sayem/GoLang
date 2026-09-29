package main

import "fmt"

func example() {
	result := 10

	defer func () {
		fmt.Println("This is from deffered function:", result)
	}()

	fmt.Println("This is from example function:", result)

	result += 100;
}

func main() {
	example()
}

/*
	Output:
		This is from example function: 10
		This is from deffered function: 110

	Note: 
		For this closure the defer function holds the "referrence" of the "result"
		variable. That's why when the value of result changed, also the deffered 
		value is changed.
*/