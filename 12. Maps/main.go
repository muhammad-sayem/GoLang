package main

import "fmt"

type user struct {
	id int 
	name string
	address string
}

func main() {
	myMap := make(map[string]int)

	myMap["user1Age"] = 25
	myMap["user2Age"] = 13
	myMap["user3Age"] = 28

	fmt.Println(myMap)
	fmt.Println(myMap["user2Age"])

	//* ====================================================== *//

	fmt.Println("--- Map Literal (Declaring and initializing values) ---");

	myMap2 := map[string]string {
		"name": "Akash",
		"address": "Dhaka",
		"maritalStatus": "Single",
	}
	fmt.Println(myMap2)
	fmt.Println(myMap2["address"])
	
	//* Deleting property from map *//
	delete(myMap2, "maritalStatus")
	fmt.Println(myMap2)

	//* ====================================================== *//

	//* Map of Struct *//
	fmt.Println("----------- Map of Struct -----------");
	myMap3 := map[string]user {
		"data": user {
			id: 1,
			name: "Srayo",
			address: "Dhaka",
		},
	}

	fmt.Println(myMap3)
}
