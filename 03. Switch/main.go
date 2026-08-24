package main

import "fmt"

func main() {
	day := "sun"
	
	//* Tagged Switch *//
	switch day {
	case "fri":
		fmt.Println("Full off day!!")
	case "sat":
		fmt.Println("Half Working Day!!")
	default:
		fmt.Println("Working Day!!")
	}

	//* Non tagged Switch *//
	// switch {
	// case day == "fri":
	// 	fmt.Println("Full off day!!")
	// case day == "sat":
	// 	fmt.Println("Half Working Day!!")
	// default:
	// 	fmt.Println("Working Day!!")
	// }

}
