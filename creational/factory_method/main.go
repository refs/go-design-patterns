package main

import "fmt"

// Kind
type Kind int

const (
	VILLAGER = iota
	METROPOLITAN
)

// Settler defines what a Settler is
type Settler interface {
	GetKind() string
}

// Villager implements the Settler interface.
type Villager struct{}

// GetName implementation of Settler.
func (v Villager) GetKind() string {
	return "villager"
}

// Metropolitan implements the Settler interface.
type Metropolitan struct{}

// GetName implementation of Settler.
func (v Metropolitan) GetKind() string {
	return "villager"
}

// GetPerson is the factory method. It returns a Settler kind according to its arguments.
func GetPerson(k Kind) (Settler, error) {
	switch k {
	case VILLAGER:
		return Villager{}, nil
	case METROPOLITAN:
		return Metropolitan{}, nil
	default:
		return nil, fmt.Errorf("invalid kind: %v", k)
	}
}

func main() {
	folk, err := GetPerson(VILLAGER)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T, I am a %s", folk, folk.GetKind()) // main.Villager, I am a villager
}
