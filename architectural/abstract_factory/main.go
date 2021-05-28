package main

import (
	"fmt"
	"runtime"
)

/*
	In normal usage, the client software creates a concrete implementation of the abstract factory and then uses the
	generic interface of the factory to create the concrete objects that are part of the theme. Let's see how that
	is reflected in this code.
*/

/*
	First of, let's define the functionality of our concrete types. Buttons draw themselves onto the screen.
*/
type button interface {
	draw()
}

/*
	Secondly, create a few concrete types that implement the behavior from the button interface:
	- windowsButton - will render only on windows
	- macButton 	- will render only on mac
	- linuxButton 	- will render only on linux
*/
type windowsButton struct{}

func (w windowsButton) draw() {
	print("[Windows button]")
}

type macButton struct{}

func (m macButton) draw() {
	print("[macOS button]")
}

type linuxButton struct{}

func (l linuxButton) draw() {
	print("[Linux button]")
}

/*
	Next let's define our Abstract Factory. This is the layer that deals with concrete types and allows the client to
	operate without any knowledge of any concrete type declared above.
*/
type buttonFactory interface {
	createButton() button
}

type buttonDrawer struct {
	factory buttonFactory
}

func factoryFromOS() (buttonFactory, error) {
	switch runtime.GOOS {
	case "darwin":
		return macFactory{}, nil
	case "windows":
		return windowsFactory{}, nil
	case "linux":
		return linuxFactory{}, nil
	default:
		return nil, fmt.Errorf("unsuported os")
	}
}

func newButtonDrawer() buttonDrawer {
	factory, err := factoryFromOS()
	if err != nil {
		panic(err)
	}

	return buttonDrawer{
		factory: factory,
	}
}

type windowsFactory struct{}

func (f windowsFactory) createButton() button {
	return windowsButton{}
}

type macFactory struct{}

func (f macFactory) createButton() button {
	return macButton{}
}

type linuxFactory struct{}

func (f linuxFactory) createButton() button {
	return linuxButton{}
}

func main() {
	client := newButtonDrawer()
	// this pattern scales horizontally: having more GUI buttons is just a matter of adding more factories and concrete
	// types. However the client logic remain untouched. Adding another set of factory and struct won't break this code.
	button := client.factory.createButton()
	button.draw()
}
