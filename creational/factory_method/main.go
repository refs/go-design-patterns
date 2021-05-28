package main

import "fmt"

// Kind
type Kind int

const (
	VILLAGER = iota
	METROPOLITAN
)

// Person defines what a Person is
type Person interface {
	GetName() string
}

// Villager implements the Person interface.
type Villager struct{}

// GetName implementation of Person.
func (v Villager) GetName() string {
	return "villager"
}

// Metropolitan implements the Person interface.
type Metropolitan struct{}

// GetName implementation of Person.
func (v Metropolitan) GetName() string {
	return "villager"
}

// GetPerson is the factory method. It returns a Person kind according to its arguments.
func GetPerson(k Kind) (Person, error) {
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

	fmt.Printf("%T, I am a %s", folk, folk.GetName()) // main.Villager
}
