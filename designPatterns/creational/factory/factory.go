package creational

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

type CashPM struct{}
type DebitCardPM struct{}

const (
	Cash      = 1
	DebitCard = 2
)

func (p *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f Paid using Cash\n", amount)
}

func (p *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using Debit Card\n", amount)
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}
