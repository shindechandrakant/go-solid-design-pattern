package main

import (
	"context"
	"fmt"
)

type ChargeProcessor interface {
	Charge(ctx context.Context, amount int64) (txnID string, err error)
}

type RefundProcessor interface {
	Refund(ctx context.Context, txnID string) error
}

type SaveCardProcessor interface {
	SaveCard(ctx context.Context, userID string, cardToken string) error
}

type FullProcessor interface {
	ChargeProcessor
	RefundProcessor
	SaveCardProcessor
}

type ChargeRefundProcessor interface {
	ChargeProcessor
	RefundProcessor
}

type CryptoProcessor struct{}

func (c *CryptoProcessor) Charge(ctx context.Context, amount int64) (string, error) {
	// calls Stripe API, creates charge
	return "crypto_txn_123", nil
}

// type PaymentProcessor interface {
//     Refund(ctx context.Context, txnID string) error
//     SaveCard(ctx context.Context, userID string, cardToken string) error
// }

type StripeProcessor struct{}

func (s *StripeProcessor) Charge(ctx context.Context, amount int64) (string, error) {
	// calls Stripe API, creates charge
	return "stripe_txn_123", nil
}

func (s *StripeProcessor) Refund(ctx context.Context, txnID string) error {
	// calls Stripe refund API
	return nil
}

func (s *StripeProcessor) SaveCard(ctx context.Context, userID string, cardToken string) error {
	// saves card to Stripe customer vault
	return nil
}

type CashOnDelivery struct{}

func (c *CashOnDelivery) Charge(ctx context.Context, amount int64) (string, error) {
	// records a pending COD order
	return "cod_pending_456", nil
}

func (c *CashOnDelivery) Refund(ctx context.Context, txnID string) error {
	// marks COD order for cash return on next visit
	return nil
}

type CheckoutService struct{}

func (s *CheckoutService) Checkout(ctx context.Context, processor ChargeProcessor, amount int64) error {
	txnID, err := processor.Charge(ctx, amount)
	if err != nil {
		return err
	}
	fmt.Println("charged:", txnID)
	return nil
}

func (s *CheckoutService) SaveUserCard(ctx context.Context, userID, token string, processor SaveCardProcessor) error {
	return processor.SaveCard(ctx, userID, token)
}
