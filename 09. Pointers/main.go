package main

import "fmt"

func changeValue (val *int) {
	*val = 200	
	fmt.Println("Inside changeValue function:", val)	// Inside changeValue function: 0x2a3194fc0120
	fmt.Println("Inside changeValue function:", *val) // Inside changeValue function: 200
}

func modifyWithoutPointer (arr [5] int) {
	arr[0] = 999
	fmt.Println("Array inside modify WITHOUT POINTER function:", arr)
}

func modifyWithPointer (arr *[5] int) {
	arr[0] = 777
	fmt.Println("Array inside modify WITH POINTER function:", *arr)
}

//* ============================================================ */

func main () {
	a := 43
	fmt.Println(a);		// 43
	fmt.Println(&a);	// 0xd80d47d00c0 (Memory Address)

	fmt.Println("--------------------------------")

	x := 50
	y := &x
	x = 100

	fmt.Println("x:", x)		// x: 100
	fmt.Println("y:", y)		// y: 0x39857217a0d8
	fmt.Println("y:", *y)		// y: 100 (Dereference: Getting the value of the pointer it holds)

	fmt.Println("----- Pointers in Function Parameters -----")

	p := 20
	changeValue(&p)
	fmt.Println("Outside changeValue function:", p)		// Outside changeValue function: 200

	fmt.Println("----- Passing Arrays as Pointer -----")
	bigArray := [5]int{10, 20, 30, 40, 50}

	modifyWithoutPointer(bigArray)
	fmt.Println("Array after modify WIHTHOUT POINTER function:", bigArray)
	
	bigArray2 := [5]int{10, 20, 30, 40, 50}

	modifyWithPointer(&bigArray2)
	fmt.Println("Array after modify WIHT POINTER function:", bigArray2)
}	