package Factory

import "fmt"

type CashPM struct {}

const CashPaymentReceipt = "Paid via cash"

func (c *CashPM) Pay(amount float64) string {
	return fmt.Sprintf("%f %s",amount, CashPaymentReceipt)
}

