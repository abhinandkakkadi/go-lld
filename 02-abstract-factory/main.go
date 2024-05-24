package main

import "fmt"

type Dog interface {
	dogSound()
}

type newDog struct{}

func (c *newDog)  dogSound() {
	fmt.Println("new dog sound")
}

type oldDog struct{}

func (c *oldDog) dogSound() {
	fmt.Println("old dog sound")
}

type Cat interface {
	catSound()
}

type newCat struct{}

func (c *newCat) catSound() {
	fmt.Println("new cat sound")
}

type oldCat struct{}

func (c *oldCat) catSound() {
	fmt.Println("old cat sound")
}

type AnimalFactory interface {
	createCat() Cat
	createDog() Dog
}

type newAnimalFactory struct{}

func (c *newAnimalFactory) createCat() Cat {
	return &newCat{}
}

func (c *newAnimalFactory) createDog() Dog {
	return &newDog{}
}

type oldAnimalFactory struct{}

func (o *oldAnimalFactory) createCat() Cat {
	return &oldCat{}
}

func (o *oldAnimalFactory) createDog() Dog {
	return &oldDog{}
}

func main() {
	factory1 := newAnimalFactory{}
	cat1 := factory1.createCat()
	dog1 := factory1.createDog()
	cat1.catSound()
	dog1.dogSound()

	factory2 := oldAnimalFactory{}
	cat2 := factory2.createCat()
	dog2 := factory2.createDog()
	cat2.catSound()
	dog2.dogSound()
}