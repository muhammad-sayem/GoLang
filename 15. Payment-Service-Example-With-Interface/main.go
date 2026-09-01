package main

import "fmt"

type PaymentMethod interface {
	pay(amount float64)
}

type Bkash struct {
	apiKey string
}

func (bk Bkash) pay(amount float64) {
	fmt.Printf("Paid %.2f by BKash\n", amount)
}

type Nagad struct {
	apiKey string
}

func (ng Nagad) pay(amount float64) {
	fmt.Printf("Paid %.2f by Nagad\n", amount)
}

//* ----------------------------------------------------- *//

type PaymentService struct {
	method PaymentMethod
}

func NewPaymentService(method PaymentMethod) PaymentService {
	return PaymentService{
		method: method,
	}
}

func (ps PaymentService) checkout(amount float64) {
	ps.method.pay(amount)
}

func main() {
	//* Pay with Bkash *//
	bkash := Bkash{apiKey: "1234abcd"}
	payWithBkash := NewPaymentService(bkash)
	payWithBkash.checkout(60.0)

	//* Pay with Nagad *//
	nagad := Nagad{apiKey: "1234abcd"}
	payWithNagad := NewPaymentService(nagad)
	payWithNagad.checkout(100.0)
}
