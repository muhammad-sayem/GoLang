package main

import "fmt"

type additionalInfo struct {
	phone string
	address string
}

type user struct {
	name  string
	email string
	extraInfo additionalInfo
}

type student struct {
	id int
	name string
	address string
}

func main() {
	// john := user{"John Doe", "john@gmail.com"} // Positional
	john := user{name: "John", email: "john@gmail.com"} // Key-value pair (mostly used)

	fmt.Printf("%+v", john) // "%+v" this is for showing key-value pair

	fmt.Println(john.name)
	fmt.Println(john.email)
	
	john.name = "John Doe"
	fmt.Println(john.name)

	var user1 user
	user1.name = "Test User 1"
	user1.email = "test1@gmail.com"
	fmt.Println(user1.name)
	fmt.Println(user1.email)

	fmt.Println("--------------- Nested Structs ---------------")
	
	newUser1 := user{
		name: "New User 1", 
		email: "newuser1@gmail.com", 
		extraInfo: additionalInfo{
			phone: "+880111111111",
			address: "Dhaka",
		},
	}
	fmt.Printf("%+v", newUser1)
	fmt.Printf("\n")

	fmt.Println("--------------- Constructor Functions ---------------")
	
	newStudent := func (id int, name string, address string) student{ 
		return student {
			id: id,
			name: name,
			address: address,
		}
	}
	
	asif := newStudent(1, "Asif", "Dhaka")
	fmt.Printf("+%v", asif)
}