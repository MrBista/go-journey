package designpattern

import (
	"fmt"
	"testing"
)

/*
3. Strategy Pattern
Strategy pattern enables selecting an algorithm’s behavior at runtime.

Use Case:
Payment gateways like PayPal, Stripe, etc.

*/

type PaymentStrategy interface {
	Pay(amount float64)
}

type PayPal struct{}

func (e *PayPal) Pay(amount float64) {
	fmt.Println("bayar by paypal dengan nominal ", amount)
}

type Stripe struct{}

func (e *Stripe) Pay(amount float64) {
	fmt.Println("bayar by string dengan nominal ", amount)
}

type PaymentContext struct {
	Strategy PaymentStrategy
}

func (pc *PaymentContext) Execute(amount float64) {
	pc.Strategy.Pay(amount)
}

func TestStrategyPattern(t *testing.T) {
	ctx := PaymentContext{Strategy: &PayPal{}}
	ctx.Execute(23)
}
