package models

import (
	"fmt"
	"time"
)

type CreditCardStrategy struct {
	Name           string
	Card           string
	ExpirationDate time.Time
	CVV            string
}

func NewCreditCardStrategy(name, card, cvv string, expDate time.Time) (result CreditCardStrategy) {
	result.Name = name
	result.Card = card
	result.CVV = cvv
	result.ExpirationDate = expDate
	return
}

func (card *CreditCardStrategy) makePay(amount float64) bool {
	fmt.Printf("Payed %f via CreditCardStrategy!\n\nCard: %s\n", amount, card.Card)
	return true
}
