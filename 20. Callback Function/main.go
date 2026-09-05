package main

import "fmt"

// func process(greet func()) {
// 	greet()
// }

func calculate(a int, b int, operation func(x int, y int) int) int {
	result := operation(a, b)
	return result
}

func main() {
	// sayHello := func() {
	// 	fmt.Println("Hello Bangladesh!!")
	// }
	// process(sayHello)

	add := func(a int, b int) int {
		return a + b
	}

	sub := func(a int, b int) int {
		return a - b
	}

	multiply := func(a int, b int) int {
		return a * b
	}

	fmt.Println(calculate(30, 10, add))
	fmt.Println(calculate(50, 40, sub))
	fmt.Println(calculate(10, 2, multiply))

	//* Annonymus Callback Function *//

	result1 := calculate(60, 30, func(x int, y int) int {
		return x + y
	})
	fmt.Println("Summation:", result1)
	
	result2 := calculate(50, 10, func(x int, y int) int {
		return x - y
	})
	fmt.Println("Subtraction:", result2)
}
