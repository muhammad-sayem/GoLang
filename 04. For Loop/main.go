package main

import "fmt"

func main() {
	//* Normal style *//
	for i := 1; i <= 5; i++ {
		fmt.Println("Count:", i)
	}

	//* While style *//
	i := 1
	for i <= 10 {
		fmt.Println("Counting...", i)
		i++
	}

	//* Break *//
	for i := 1; i <= 20; i++ {
		if i == 8 {
			break
		}
		fmt.Println("Break checking...", i)
	}

	//* Continue *//
	for i := 1; i <= 20; i++ {
		if(i % 2 != 0){
			continue
		}
		fmt.Println("Continue checking...", i)
	}
}
