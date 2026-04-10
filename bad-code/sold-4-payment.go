package bad_code

import (
	"fmt"
	"time"
)

var dbConnection = "postgres://localhost:5432/payments"
var stripeKey = "sk_live_xxx"
var paypalSecret = "pp_secret_xxx"

type PaymentService struct {
	logs []string
}

func (p *PaymentService) ProcessPayment(
	userId int,
	amount float64,
	currency string,
	paymentMethod string,
	cardNumber string,
	paypalEmail string,
	cryptoWallet string,
	couponCode string,
	isSubscription bool,
	subscriptionInterval string,
) (string, error) {

	// logging
	p.logs = append(p.logs, fmt.Sprintf("[%s] Payment started for user %d", time.Now(), userId))

	// validation
	if amount <= 0 {
		return "", fmt.Errorf("invalid amount")
	}
	if currency != "USD" && currency != "EUR" && currency != "GBP" {
		return "", fmt.Errorf("unsupported currency")
	}

	// apply coupon
	finalAmount := amount
	if couponCode == "SAVE10" {
		finalAmount = amount * 0.9
	} else if couponCode == "SAVE20" {
		finalAmount = amount * 0.8
	} else if couponCode == "HALFOFF" {
		finalAmount = amount * 0.5
	} else if couponCode != "" {
		return "", fmt.Errorf("invalid coupon")
	}

	// currency conversion
	if currency == "EUR" {
		finalAmount = finalAmount * 1.1
	} else if currency == "GBP" {
		finalAmount = finalAmount * 1.3
	}

	// process payment
	var transactionId string
	if paymentMethod == "card" {
		fmt.Printf("Connecting to Stripe with key: %s\n", stripeKey)
		fmt.Printf("Charging card %s for %.2f\n", cardNumber, finalAmount)
		transactionId = fmt.Sprintf("stripe_%d", time.Now().UnixNano())
		// retry logic for card
		for i := 0; i < 3; i++ {
			fmt.Println("Attempting card charge...")
			break // pretend it worked
		}
	} else if paymentMethod == "paypal" {
		fmt.Printf("Using PayPal secret: %s\n", paypalSecret)
		fmt.Printf("Charging PayPal %s for %.2f\n", paypalEmail, finalAmount)
		transactionId = fmt.Sprintf("paypal_%d", time.Now().UnixNano())
	} else if paymentMethod == "crypto" {
		fmt.Printf("Sending %.2f to wallet %s\n", finalAmount, cryptoWallet)
		transactionId = fmt.Sprintf("crypto_%d", time.Now().UnixNano())
	} else {
		return "", fmt.Errorf("unsupported payment method")
	}

	// handle subscription
	if isSubscription {
		if subscriptionInterval == "monthly" {
			fmt.Println("Setting up monthly billing")
		} else if subscriptionInterval == "yearly" {
			fmt.Println("Setting up yearly billing with 10% discount")
			finalAmount = finalAmount * 0.9
		} else {
			return "", fmt.Errorf("invalid subscription interval")
		}
		fmt.Printf("Subscription created for user %d\n", userId)
	}

	// save to database
	fmt.Printf("Saving to DB: %s\n", dbConnection)
	fmt.Printf("Transaction %s: user=%d, amount=%.2f\n", transactionId, userId, finalAmount)

	// send notification
	fmt.Printf("Sending email receipt to user %d\n", userId)
	if finalAmount > 1000 {
		fmt.Println("High value transaction - sending SMS alert")
	}

	// logging
	p.logs = append(p.logs, fmt.Sprintf("[%s] Payment completed: %s", time.Now(), transactionId))

	return transactionId, nil
}

func (p *PaymentService) GetLogs() []string {
	return p.logs
}

func (p *PaymentService) ValidateCard(cardNumber string) bool {
	return len(cardNumber) == 16
}

func (p *PaymentService) ConvertCurrency(amount float64, from string, to string) float64 {
	// messy conversion logic
	return amount
}

func main() {
	service := &PaymentService{}
	txId, err := service.ProcessPayment(
		1,
		99.99,
		"USD",
		"card",
		"4111111111111111",
		"",
		"",
		"SAVE10",
		false,
		"",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction:", txId)
}
