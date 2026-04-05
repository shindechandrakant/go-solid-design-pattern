package bad_code

import (
	"fmt"
)

type OrderService struct{}

func (o *OrderService) PlaceOrder(userType string, paymentType string, amount float64) {

	// validation
	if amount <= 0 {
		panic("invalid amount")
	}

	// pricing logic
	finalAmount := amount
	if userType == "PREMIUM" {
		finalAmount = amount * 0.9
	}

	// payment logic
	if paymentType == "CARD" {
		fmt.Println("Processing card payment:", finalAmount)
	} else if paymentType == "UPI" {
		fmt.Println("Processing UPI payment:", finalAmount)
	} else {
		panic("unsupported payment type")
	}

	// persistence
	o.saveToDB(finalAmount)

	// notification
	if userType == "PREMIUM" {
		fmt.Println("Sending WhatsApp notification")
	} else {
		fmt.Println("Sending Email notification")
	}
}

func (o *OrderService) saveToDB(amount float64) {
	fmt.Println("Saving order:", amount)
}

func main() {
	service := OrderService{}
	service.PlaceOrder("PREMIUM", "CARD", 1000)
}
