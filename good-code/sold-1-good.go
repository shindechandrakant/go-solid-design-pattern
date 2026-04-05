package good_code

import "fmt"

type OrderService struct {
	paymentProcessor PaymentProcessor
	repository       OrderRepository
	strategy         PricingStrategy
}

type PaymentType string

const (
	UPI  PaymentType = "UPI"
	CARD PaymentType = "CARD"
)

type PaymentProcessor interface {
	Process(amount float64)
}

type CardPaymentMethod struct{}
type UPIPaymentMethod struct{}

func (cpm *CardPaymentMethod) Process(amount float64) {
	fmt.Printf("Processing %f using Card", amount)
}

func (upm *UPIPaymentMethod) Process(amount float64) {
	fmt.Printf("Processing %f using UPI", amount)
}

type OrderRepository interface {
	save(amount float64)
}

type CreateOrder struct{}

func (co *CreateOrder) save(amount float64) {
	fmt.Printf("Saving order: %f", amount)
}

type Notifier interface {
	Notify()
}

type EmailNotification1 struct {
}

func (e *EmailNotification) Notify() {
	fmt.Println("sending notification")
}

type PricingStrategy interface {
	Calculate(float64) float64
}

type NormalUser struct {
	notifier Notifier
}
type PremiumUser struct {
	notifier Notifier
}

func (pu *PremiumUser) Calculate(amount float64) float64 {
	return amount * 0.9
}

func (pu *NormalUser) Calculate(amount float64) float64 {
	return amount
}

func (o *OrderService) PlaceOrder(amount float64) {

	if amount <= 0 {
		panic("In valid amount")
	}

	finalAmount := o.strategy.Calculate(amount)
	o.paymentProcessor.Process(finalAmount)
	o.repository.save(finalAmount)

}

func (o *OrderService) SaveToDB(amount float64) {
	fmt.Println("Saving order: ", amount)
}

var PaymentRegistry = map[PaymentType]func() PaymentProcessor{
	UPI: func() PaymentProcessor {
		return &UPIPaymentMethod{}
	},
	CARD: func() PaymentProcessor {
		return &CardPaymentMethod{}
	},
}

func PaymentMethodFactory(pt PaymentType) PaymentProcessor {
	//switch pt {
	//case UPI:
	//	return &UPIPaymentMethod{}
	//case CARD:
	//	return &CardPaymentMethod{}
	//}
	if creator, ok := PaymentRegistry[pt]; ok {
		return creator()
	}
	panic("In valid payment type")
}

func NewOrderService(processor PaymentProcessor, repo OrderRepository) *OrderService {

	return &OrderService{
		paymentProcessor: processor,
		repository:       repo,
	}
}

func main() {
	//userType := PremiumUser{
	//	notifier: &EmailNotification{},
	//}
	upi := PaymentMethodFactory(UPI)
	order := NewOrderService(upi, &CreateOrder{})

	order.PlaceOrder(100)
}
