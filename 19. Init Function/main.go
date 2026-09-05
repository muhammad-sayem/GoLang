package main

import "fmt"

//* init() function executes before main() function *//

func init () {
	fmt.Println("I'm from init function...")
}

func main() {
	fmt.Println("I'm from main function...")
}