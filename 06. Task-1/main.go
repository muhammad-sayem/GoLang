package main

import "fmt"

func calculateGrade(marks int) string {
	switch {
	case marks < 0 || marks > 100:
		return "Invalid input"
	case marks >= 90 && marks <= 100:
		return "A+"
	case marks >= 80 && marks < 90:
		return "A"
	case marks >= 70 && marks < 80:
		return "B"
	case marks >= 60 && marks < 70:
		return "C"
	default:
		return "Fail"
	}
}

func passFailStatus(marks int) {
	if marks >= 60 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func main() {

	displayMenu := func() {
		fmt.Println("Welcome to grade calculator!!")
		fmt.Println("1) Calculate Grade")
		fmt.Println("2) Check Pass/Fail Status")
		fmt.Println("3) Exit the program")
		fmt.Println("Switch an option")
	}

	var choice int
	var marks int

	running := true
	for running {
		displayMenu()

		fmt.Printf("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Enter marks: ")
			fmt.Scan(&marks)

			result := calculateGrade(marks)
			if result == "Invalid input" {
				fmt.Println("Choose a valid number between 1 to 100...")
			} else {
				fmt.Println("You got:", result)
			}

		case 2:
			fmt.Printf("Enter marks: ")
			fmt.Scan(&marks)

			passFailStatus(marks)

		case 3:
			fmt.Println("Exiting Program...")
			running = false

		default:
			fmt.Println("Choose a valid input from menu")
		}
	}
}
