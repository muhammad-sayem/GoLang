package main

import "fmt"

func main() {
	var age int = 13

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Not adult")
	}

	score := 70

	if score >= 80 {
		fmt.Println("A+")
	} else if score >= 70 && score < 80 {
		fmt.Println("A")
	} else if score >= 60 && score < 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	//* Classic Style *//

	if money := 150; money < 100 {
		fmt.Println("Buy Chocolates")
	} else if money >= 100 && money < 200 {
		fmt.Println("Buy French Fry")
	} else if money >= 200 && money < 300 {
		fmt.Println("Buy Burger")
	} else if money >= 300 {
		fmt.Println("Buy Combo of French Fry and Burger")
	}
}
