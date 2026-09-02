package main

import "fmt"

//* interface {} == any *//

func print(data any) {
	fmt.Println("Data:", data)
}

func process(data any) {
	strData, ok := data.(string) // assuming that the data is string type and "ok" returns boolean
	if ok {
		fmt.Println("Length of the data:", len(strData))
	} 

	intData, ok := data.(int)		// // assuming that the data is integer type and "ok" returns boolean
	if ok {
		fmt.Println("The number is:", intData)
	}
}

func main() {
	print([]int{4, 8, 5})
	print("Muhammad Sayem")

	process("Masud Abdullah")
	process(10)
}
