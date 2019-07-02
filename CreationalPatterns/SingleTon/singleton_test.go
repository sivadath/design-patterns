package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter := GetInstance()

	//Test of acceptance criteria 1, Call to GetInstance should never return nil
	if counter == nil {
		t.Error("Expected Pointer to Singleton object , but received nil")
	}
	currentCounter := counter.Increment()

	if currentCounter != 1 {
		t.Errorf("After the first call to increment count should be equal to 1 but it is %d", currentCounter)
	}
}
