package creational

import (
	"errors"
	"fmt"
)

type PaymentType int

const (
	Cash PaymentType = iota
	DebitCard
)

type PaymentMethod interface {
	Pay() string
}

type CashMethod struct{}
type DebitMethod struct{}

func (c CashMethod) Pay() string {
	return "Paid with cash"
}

func (d DebitMethod) Pay() string {
	return "Paid with Debit"
}

func GetPaymentMethod(method PaymentType) (PaymentMethod, error) {
	switch method {
	case Cash:
		return CashMethod{}, nil
	case DebitCard:
		return DebitMethod{}, nil
	default:
		fmt.Printf("Invalid payment type supplied")
		return DebitMethod{}, errors.New("invalid payment")
	}
}
