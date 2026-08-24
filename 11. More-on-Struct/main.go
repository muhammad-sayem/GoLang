package main

import "fmt"

type user struct {
	name string
	age int
	isLoggedIn bool
}

func main() {
	fmt.Println("---------- Using Receiver Function ----------");

	user1 := user {
		name: "Masud",
		age: 24,
		isLoggedIn: false,
	}

	user1.greet();

	//* ================================================== *//

	fmt.Println("------ Receiver Function (With Pointer) ------");

	// user1.login() 			// It will not work
	// pointerUser1 := &user1 		// Go handles this internally 
	user1.login()
	fmt.Printf("%+v", user1)
}

//* ======================================================== *//

// * Receiver Function *//
func (u user) greet () {
	fmt.Println("Hello", u.name)
}

func (u *user) login () {
	fmt.Println("Login Called...")
	// (*u).isLoggedIn = true
	u.isLoggedIn = true
}