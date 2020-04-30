package ProtoType

import "testing"

func TestClone(t *testing.T)  {
	shirtsCache := GetShirtsCloner()

	if shirtsCache == nil {
		t.Fatal("Received nil shirt cache from cloner function")
	}

	item1, err := shirtsCache.GetClone(White)

	if err != nil {
		t.Error(err)
	}

	if item1 == whitePrototype {

	}
}