//* Used with array, slice, string, map, channel *//

package main

import "fmt"

func main() {
	//* With Map *//
	fmt.Println("---------- With Map ----------")
	myMap := map[string]string {
		"name": "Masud",
		"address": "Dhaka",
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	//* With Array *//
	fmt.Println("---------- With Array ----------")
	colors := [3]string {
		"White",
		"Black",
		"Red",
	}

	for idx, value := range colors {
		fmt.Println(idx, value)
	}

	//* With Slice *//
	fmt.Println("---------- With Slice ----------")
	
	animals := []string {
		"Horse",
		"Cow",
		"Tiger",
		"Lion",
	}

	for idx, value := range animals {
		fmt.Println(idx, value)
	}

	//* With String *//
	fmt.Println("---------- With String ----------")

	name := "Shihab"
	for idx, value := range name {
		fmt.Println(idx, value)		// returns ASCII values of characters
	}
}
