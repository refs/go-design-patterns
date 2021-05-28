package main

import "fmt"

// Abstract Factories

// IClothesFactory makes clothes
type IClothesFactory interface {
	produceShirt() IShirt
	produceHat() IHat
}

func getClothesFactory(name string) IClothesFactory {
	switch name {
	case "nike":
		return nike{}
	case "adidas":
		return adidas{}
	default:
		panic("unsuported factory")
	}
}

// Concrete Factories

// nike implements the IClothesFactory
type nike struct{}

func (n nike) produceShirt() IShirt {
	s := shirt{
		material: "cotton",
		design:   "nike-wild-zebra",
	}

	return &s
}

func (n nike) produceHat() IHat {
	h := hat{
		size:    "m",
		hatType: "nike-cap",
	}

	return &h
}

// adidas implements the IClothesFactory
type adidas struct{}

func (a adidas) produceShirt() IShirt {
	s := shirt{
		material: "cotton",
		design:   "nike-wild-zebra",
	}

	return &s
}

func (a adidas) produceHat() IHat {
	h := hat{
		size:    "m",
		hatType: "adidas-cap",
	}

	return &h
}

// Abstract Products

type IShirt interface {
	SetMaterial(string)
	SetDesign(string)
	GetMaterial() string
	GetDesign() string
}

type IHat interface {
	SetSize(string)
	SetHatType(string)
	GetSize() string
	GetHatType() string
}

// Concrete Products

type shirt struct {
	material string
	design   string
}

func (s *shirt) SetMaterial(m string) {
	s.material = m
}

func (s *shirt) SetDesign(d string) {
	s.design = d
}

func (s *shirt) GetMaterial() string {
	return s.material
}

func (s *shirt) GetDesign() string {
	return s.design
}

type hat struct {
	size    string
	hatType string
}

func (h *hat) SetSize(s string) {
	h.size = s
}

func (h *hat) SetHatType(s string) {
	h.hatType = s
}

func (h *hat) GetSize() string {
	return h.size
}

func (h *hat) GetHatType() string {
	return h.hatType
}

// Client code
func main() {
	nike := getClothesFactory("nike")
	nikeHat := nike.produceHat()
	fmt.Println(nikeHat)

	adidas := getClothesFactory("adidas")
	adidasHat := adidas.produceHat()
	fmt.Println(adidasHat)
}
