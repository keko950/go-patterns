package creational

import (
	"fmt"
	"testing"
)

func TestPaymentMethodDoesntExist(t *testing.T) {
	_, err := GetPaymentMethod(3)

	if err == nil {
		t.Errorf("payment method 3 should not exist")
	}
}

func TestPaymentMethodCash(t *testing.T) {
	pay, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Errorf("payment method should be cash")
	}

	fmt.Println(pay.Pay())
}

func TestPaymentMethodDebitCard(t *testing.T) {
	pay, err := GetPaymentMethod(DebitCard)

	if err != nil {
		t.Errorf("payment method should be debit card")
	}

	fmt.Println(pay.Pay())
}
