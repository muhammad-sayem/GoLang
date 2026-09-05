package main //* Executable package *//

import (
	"learngo/payment"
	"github.com/fatih/color"
)

func main() {
	//* Pay with Bkash *//
	bkash := payment.Bkash{ApiKey: "1234abcd"}
	payWithBkash := payment.NewPaymentService(bkash)
	payWithBkash.Checkout(60.0)

	//* Pay with Nagad *//
	nagad := payment.Nagad{ApiKey: "1234abcd"}
	payWithNagad := payment.NewPaymentService(nagad)
	payWithNagad.Checkout(100.0)

	color.Green("The color is Green")
	color.BgRGB(230, 42, 42).Println("Background Red")
}
