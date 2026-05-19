package good_code

//payment methods Types -> strategy
//1. Card ->  retry logic
//2. paypal
//3. Crypto

//Rules
// currency converters
// coupon codes

type PaymentType string
type Currency string

const (
	PAYPAL PaymentType = "PAYPAL"
	CARD   PaymentType = "CARD"
	CRYPTO PaymentType = "CRYPTO"
)

type PaymentProcessor interface {
	Process(ctx PaymentContext) (string, error)
}

type PaymentContext struct {
	Amount   float64
	Currency Currency
}

type Card struct {
	CardNumber string
	StripeKey  string
}
type Paypal struct {
	Email        string
	PaypalSecret string
}
type Crypto struct {
	WalletAddress string
}

func (c *Card) Process(ctx PaymentContext) (string, error) {
	return "TX", nil
}

func (c *Paypal) Process(ctx PaymentContext) (string, error) {
	return "TX", nil
}

func (c *Crypto) Process(ctx PaymentContext) (string, error) {
	return "TX", nil
}

var PaymentRegistry = map[PaymentType]func() PaymentProcessor{
	PAYPAL: func() PaymentProcessor {
		return &Paypal{}
	},
	CRYPTO: func() PaymentProcessor {
		return &Crypto{}
	},
	CARD: func() PaymentProcessor {
		return &Card{}
	},
}
