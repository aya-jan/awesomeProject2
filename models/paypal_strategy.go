package models

import (
	"fmt"
)

type PayPalStrategy struct {
	Email    string
	Password string
}

func NewPayPalStrategy(email, password string) (result PayPalStrategy) {
	result.Email = email
	result.Password = password
	return
}

func (card *PayPalStrategy) makePay(amount float64) bool {
	fmt.Printf("Payed %f via PayPal!\n\nEmail: %s\nPassword: %s", amount, card.Email, card.Password)
	return true
}
