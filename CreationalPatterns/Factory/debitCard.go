package Factory

import "fmt"

type DebitCardPM struct {}

const DebitCardPaymentReceipt = "Paid via debit card"

func (d *DebitCardPM) Pay(amount float64) string {
	return fmt.Sprintf("%f %s",amount, DebitCardPaymentReceipt)
}

