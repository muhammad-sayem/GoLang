package payment

import "fmt"

type paymentMethod interface {
	Pay(amount float64)
}

type Bkash struct {
	ApiKey string
}

func (bk Bkash) Pay(amount float64) {
	fmt.Printf("Paid %.2f by BKash\n", amount)
}

type Nagad struct {
	ApiKey string
}

func (ng Nagad) Pay(amount float64) {
	fmt.Printf("Paid %.2f by Nagad\n", amount)
}

//* ----------------------------------------------------- *//

type PaymentService struct {
	method paymentMethod
}

func NewPaymentService(method paymentMethod) PaymentService {
	return PaymentService{
		method: method,
	}
}

func (ps PaymentService) Checkout(amount float64) {
	ps.method.Pay(amount)
}