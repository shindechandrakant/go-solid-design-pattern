package best_code

import (
	"errors"
	"fmt"
)

//
// ---- Payment Processor ----
//

type PaymentProcessor interface {
	Process(amount float64) error
}

type CardPayment struct{}

func (c *CardPayment) Process(amount float64) error {
	fmt.Printf("Processing %.2f using Card\n", amount)
	return nil
}

type UPIPayment struct{}

func (u *UPIPayment) Process(amount float64) error {
	fmt.Printf("Processing %.2f using UPI\n", amount)
	return nil
}

//
// ---- Payment Factory (OCP compliant) ----
//

type PaymentType string

const (
	UPI  PaymentType = "UPI"
	CARD PaymentType = "CARD"
)

var paymentRegistry = map[PaymentType]func() PaymentProcessor{
	UPI:  func() PaymentProcessor { return &UPIPayment{} },
	CARD: func() PaymentProcessor { return &CardPayment{} },
}

func PaymentFactory(pt PaymentType) (PaymentProcessor, error) {
	if creator, ok := paymentRegistry[pt]; ok {
		return creator(), nil
	}
	return nil, errors.New("invalid payment type")
}

//
// ---- Pricing Strategy ----
//

type PricingStrategy interface {
	Calculate(amount float64) float64
}

type NormalPricing struct{}

func (n *NormalPricing) Calculate(amount float64) float64 {
	return amount
}

type PremiumPricing struct{}

func (p *PremiumPricing) Calculate(amount float64) float64 {
	return amount * 0.9
}

//
// ---- Repository ----
//

type OrderRepository interface {
	Save(amount float64) error
}

type InMemoryOrderRepo struct{}

func (r *InMemoryOrderRepo) Save(amount float64) error {
	fmt.Printf("Saving order: %.2f\n", amount)
	return nil
}

//
// ---- Notifier ----
//

type Notifier interface {
	Notify(message string) error
}

type EmailNotifier struct{}

func (e *EmailNotifier) Notify(message string) error {
	fmt.Println("Sending Email:", message)
	return nil
}

//
// ---- Order Service ----
//

type OrderService struct {
	paymentProcessor PaymentProcessor
	repository       OrderRepository
	pricingStrategy  PricingStrategy
	notifier         Notifier
}

func NewOrderService(
	p PaymentProcessor,
	r OrderRepository,
	s PricingStrategy,
	n Notifier,
) *OrderService {
	return &OrderService{
		paymentProcessor: p,
		repository:       r,
		pricingStrategy:  s,
		notifier:         n,
	}
}

func (o *OrderService) PlaceOrder(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}

	finalAmount := o.pricingStrategy.Calculate(amount)

	if err := o.paymentProcessor.Process(finalAmount); err != nil {
		return err
	}

	if err := o.repository.Save(finalAmount); err != nil {
		return err
	}

	if err := o.notifier.Notify("Order placed successfully"); err != nil {
		return err
	}

	return nil
}

//
// ---- Main ----
//

func main() {

	// choose dependencies
	payment, err := PaymentFactory(UPI)
	if err != nil {
		panic(err)
	}

	repo := &InMemoryOrderRepo{}
	pricing := &PremiumPricing{}
	notifier := &EmailNotifier{}

	// inject dependencies
	orderService := NewOrderService(payment, repo, pricing, notifier)

	// execute
	err = orderService.PlaceOrder(1000)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
