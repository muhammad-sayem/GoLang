package main

import "fmt"

//* By using for loop *//
// func addNumbers(numbers ...int) int {
// 	total := 0

// 	for i := 0; i < len(numbers); i++ {
// 		total += numbers[i]
// 	}

// 	return total
// }


//* By using for range *//
func addNumbers (numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	
	return total
}

func greet (prefix string, persons ...string) {
	for _, person := range persons {
		fmt.Println(prefix, person)
	}
}

func main() {
	sum := addNumbers(10, 20, 15, 25)
	fmt.Println("Sum:",sum)

	persons := []string {"Asif", "Masud", "Srayo", "Akash", "Abir", "Sourov"}
	greet("Wlecome", persons ...) // Sending slice "persons" as argument
}
