package main

import "fmt"

func defferedFunction (result int) {
	fmt.Println("Result from deffered function:", result*2)
}

func example() {
	result := 10;

	defer defferedFunction(result)

	fmt.Println("Result from example function:", result)
}

func main() {
	defer fmt.Println("I am from deffered print call")	//* Defferd Function*//
	fmt.Println("I am from main function")

	example();
}