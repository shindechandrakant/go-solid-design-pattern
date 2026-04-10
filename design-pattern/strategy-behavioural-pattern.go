package design_pattern

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) bool
}

type CreditCardPayment struct {
	cardNumber string
}

func NewCreditCardPayment(cardNumber string) *CreditCardPayment {
	return &CreditCardPayment{
		cardNumber: cardNumber,
	}
}

func (c *CreditCardPayment) Pay(amount float64) bool {
	return true
}

type PayPalPayment struct {
	email string
}

func NewPayPalPayment(email string) *PayPalPayment {
	return &PayPalPayment{email: email}
}

func (p *PayPalPayment) Pay(amount float64) bool {
	fmt.Printf("Paid %.2f with PayPal account %s\n", amount, p.email)
	return true
}

type ShoppingCart struct {
	paymentStrategy PaymentStrategy
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (c *ShoppingCart) SetPaymentStrategy(paymentStrategy PaymentStrategy) {
	c.paymentStrategy = paymentStrategy
}

func (c *ShoppingCart) Checkout(amount float64) {
	if c.paymentStrategy != nil {
		c.paymentStrategy.Pay(amount)
	}
}

// usage
func main() {
	cart := NewShoppingCart()
	cart.SetPaymentStrategy(NewCreditCardPayment("1233"))
	cart.Checkout(123)
	cart.SetPaymentStrategy(NewPayPalPayment("ok@gmail.com"))
	cart.Checkout(234)
}
