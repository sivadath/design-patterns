package ProtoType

import "errors"

const (
	White  = iota
	Black
	Blue
)

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:"Empty",
	Color:White,
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (S *Shirt) GetInfo() string {
	return ""
}

func (S *Shirt) GetPrice() float32 {
	return S.Price
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

func GetShirtsCloner() ShirtCloner {
	return nil
}

type ShirtsCache struct {

}

func (SC *ShirtsCache) GetClone(s int) (ItemInfoGetter, error) {
	return nil, errors.New("Not yet implemented")
}



