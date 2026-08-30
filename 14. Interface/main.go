package main

import "fmt"

//* TS => Interface is used for making shape of data *//
//* GoLang => Interface is used for behavior contact *//

//* Interface *//
type Animal interface {
	sound()
}

//* For Dog *//
type Dog struct{}

func (d Dog) sound() {
	fmt.Println("Woof Woof!!")
}

//* For Cat *//
type Cat struct{}

func (c Cat) sound() {
	fmt.Println("Meow Meow!!")
}

//* For Human *//
type Human struct{
	name string
}

func (h Human) sound() {
	fmt.Println("My name is", h.name)
}

func makeSound (x Animal) {
	x.sound()
}

//* ============================================================== *//

func main() {
	tommy := Dog{}
	kitty := Cat{}
	masud := Human{"Masud Abdullah"}

	makeSound(tommy);
	makeSound(kitty);
	makeSound(masud);
}
