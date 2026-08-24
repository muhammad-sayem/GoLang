package main

import "fmt";

func makeCoffee (coffeeType string) {
	fmt.Printf("Making %s coffee..... \n", coffeeType)
}

func add (num1 int, num2 int) int {
	total := num1 + num2
	return total;
}

func buyProduct (product string, quantity int, price int) (string, int){
	totalPrice := price*quantity;
	return product , totalPrice;
}

func main() {
	// var name string = "Sayem";
	// var name = "Muhammad Sayem";

	//* Short hand way (Only allowed to use inside a function) *//
	// name := "Sayem"

	//* Grouped Variable Declaration *//
	var (
		name string = "Sayem"
		age  int    = 24
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Printf("My name is %s and age is %d \n", name, age)

	//* Multiple Variables Declaration (Allowed only for same data types) *//
	var a, b int = 30, 40;
	var x, y string = "Sayem", "Asif";

	fmt.Println(a, b)
	fmt.Println(x, y)

	//* Constant *//
	const pi = 3.1416
	fmt.Println(pi)

	//* Precision *//
	result := 4.56578
	fmt.Printf("The result is %.2f \n", result);

	//* Function Calling *//
	makeCoffee("Black");
	makeCoffee("Cold");

	addResult := add(5, 7)
	fmt.Printf("Add Result: %d \n", addResult)
	// fmt.Printf("Add Result: %d", add(8, 9));

	product, totalBill := buyProduct("Laptop", 2, 80000);
	fmt.Printf("The product is %s and total bill: %d", product, totalBill)
}
