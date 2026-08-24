package main

import "fmt"

func main() {
	//* Making Slice based on an array *//
	orders := [6]int{10, 20, 30, 40, 50, 60}

	slice1 := orders[1:4]	// [start: end (excluded)]
	fmt.Println("Slice 1...",slice1)		// [20 30 40] 

	slice2 := orders[:5]	// From first element to chosen index
	fmt.Println("Slice 2...",slice2)		// [10 20 30 40 50]
	
	slice3 := orders[2:]	// From starting (chosen) index to last element
	fmt.Println("Slice 3...",slice3)		// [30 40 50 60]

	slice4 := orders[:]	  // full array copy
	fmt.Println("Slice 4...",slice4)		// [10 20 30 40 50 60]

	fmt.Println("-------------------------------")

	//* Length and Capacity *//
	fmt.Println("The length of slice1 is:", len(slice1));
	fmt.Println("The capacity of slice1 is:", cap(slice1));

	slice1 = append(slice1, 500)
	slice1 = append(slice1, 600)

	fmt.Println("The length of slice1 is:", len(slice1));
	fmt.Println("The capacity of slice1 is:", cap(slice1));

	fmt.Println("-------------------------------")
	fmt.Println("Slice: ", slice1)
	fmt.Println("Orders: ", orders)

	slice1 = append(slice1, 700)
	slice1 = append(slice1, 800)
	
	fmt.Println("The length of slice1 is:", len(slice1));
	fmt.Println("The capacity of slice1 is:", cap(slice1));

	fmt.Println("-------------------------------")
	fmt.Println("Slice: ", slice1)
	fmt.Println("Orders: ", orders)

	fmt.Println("--------- Slice Without Array ---------")

	//* Making Slice without any array *//
	var newSlice = []int{1, 2, 3}
	fmt.Println(newSlice)
	fmt.Println("Capacity:", cap(newSlice))

	newSlice = append(newSlice, 100)
	fmt.Println(newSlice)
	fmt.Println("Capacity:", cap(newSlice))

}