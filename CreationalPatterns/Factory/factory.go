package Factory

import "errors"

const (
	Cash  = iota
	DebitCard
)

func GetPaymentMethod(method int) (PaymentMethod, error) {
	switch method {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New("unknown payment method received")

	}
	return nil, errors.New("not implemented yet")
}

type PaymentMethod interface {
	Pay(amount float64) string
}


