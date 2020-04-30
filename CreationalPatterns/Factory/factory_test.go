package Factory

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	pm, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Error("Failed to create cash payment method",err.Error())
		return
	}
	receipt := pm.Pay(10.01)
	if !strings.Contains(receipt, CashPaymentReceipt) {
		t.Error("In correct cash payment receipt received")
	}
	t.Log("Receipt received:",receipt)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	pm, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Error("Failed to create debit card payment method",err.Error())
		return
	}
	receipt := pm.Pay(10.01)
	if !strings.Contains(receipt,DebitCardPaymentReceipt) {
		t.Error("In correct debit card payment receipt received")
	}
	t.Log("Receipt received:",receipt)
}

func TestGetPaymentUnknownPaymentMethod(t *testing.T) {
	pm, err := GetPaymentMethod(10)
	if err == nil {
		t.Error("GetPaymentMethod is execpted to return error for input 10")
		return
	}
	t.Log("Received payment method:",pm)
	t.Log("Received error:",err.Error())
}