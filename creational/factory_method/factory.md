# Table of Contents

- [Factory Method](#factory-method)
  * [Code Sample Breakdown](#code-sample-breakdown)
    + [Abstract Interface](#abstract-interface)
    + [Concrete Interface](#concrete-interface)
    + [Factory Method](#factory-method-1)
    + [Supporting logic](#supporting-logic)
    + [Putting it all together: The Client](#putting-it-all-together-the-client)
  * [Recipe](#recipe)

# Factory Method

> Define an interface for creating an object, but let subclasses decide which class to instantiate

<details>
<summary>diagram</summary>
<br>

```
┌──────────────┐                      ┌──────────────┐             
│              │                      │              │             
│    Client    │─────────────────────▶│   Factory    │             
│              │                      │              │             
└──────────────┘                      └──────────────┘             
                                              △
                                              │
                                              │                   
                                      ┌───────────────┐            
                                      │               │            
                                      │    Settler    │            
                                      │               │            
                                      └───────────────┘            
                                              △
                                              │
                                  ┌ ─ ─ ─ ─ ─ ┴ ─ ─ ─ ─ ─ ┐        
                                                                   
                                  │                       │        
                          ┌───────────────┐       ┌───────────────┐
                          │               │       │               │
                          │   Villager    │       │ Metropolitan  │
                          │               │       │               │
                          └───────────────┘       └───────────────┘
```

</details>

Breaking down the definition and having a look at the code sample we have:

1. Define an interface for creating an object
2. Let subclasses decide which class to instantiate

In our example the interface we define to act as a factory is `GetPerson`.

> Note: We refer to "Interface" not as a Go interface type but rather as something we interact with, as in: "a point where two systems, subjects, organizations, etc. meet and interact".

The interaction is done is by selecting the kind of settler we initialize:

```go
folk, err := GetPerson(VILLAGER)
```

## Code Sample Breakdown

### Abstract Interface
First of, define the interface we want to abstract and ultimate interact with from the client.

```go
type Settler interface {
	GetKind() string
}
```

### Concrete Interface
Create structs that satisfy our abstract interface: `Villager` and `Metropolitan`

```go
type Villager struct{}

func (v Villager) GetKind() string {
	return "villager"
}
```

```go
type Metropolitan struct{}

func (v Metropolitan) GetKind() string {
	return "villager"
}
```

### Factory Method
Implement our factory method. Notice that we return an interface value, not the concrete type.

```go
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
```

### Supporting logic

In order to make use of some languages features, we make use `iota`s instead of strings to select the struct we interact with.

```go
type Kind int

const (
	VILLAGER = iota
	METROPOLITAN
)
```

### Putting it all together: The Client
The client point of view is just a matter of calling the factory method whenever a `Settler` type is needed:

```go
func main() {
	folk, err := GetPerson(VILLAGER)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T, I am a %s", folk, folk.GetKind()) // main.Villager, I am a villager
}
```

## Recipe

A ready-to-roll basic recipe.

```go
package main

import (
	"fmt"
)

// AbstractInterface is your abstract interface
type AbstractInterface interface {
	SomeBehavior()
}

// concrete interface types
type aConcreteTypeA struct{}

// SomeBehavior satisfies AbstractInterface
func (a aConcreteTypeA) SomeBehavior() {
	print("void-a")
}

// concrete interface types
type aConcreteTypeB struct{}

// SomeBehavior satisfies AbstractInterface
func (b aConcreteTypeB) SomeBehavior() {
	print("void-b")
}

// GetConcreteInterfaceType
func GetConcreteInterfaceType(t string) (AbstractInterface, error) {
	switch t {
	case "a":
		return aConcreteTypeA{}, nil
	case "b":
		return aConcreteTypeB{}, nil
	default:
		return nil, fmt.Errorf("invalid type: %v", t)
	}
}

// example client
func main() {
	t, err := GetConcreteInterfaceType("a")
	if err != nil {
		panic(err)
	}

	t.SomeBehavior()
}
```