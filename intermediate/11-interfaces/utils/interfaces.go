package utils

import "fmt"

type Animal interface {
	Sound() string
	Move() string
}

type Dog struct {
	Name string
}

func (d Dog) Sound() string {
	return "Woof !"
}

func (d Dog) Move() string {
	return "Running !"
}

type Cat struct {
	Name string
}

func (c Cat) Sound() string {
	return "Mewoo"
}

func (d Cat) Move() string {
	return "Jumping"
}

func AnimalBehavior(a Animal) {
	fmt.Printf("Sound animal: %s, Movement: %s\n", a.Sound(), a.Move())
}

func InterfaceLearn() {
	fmt.Println("Memanggil interface untuk dipelajari")
	dog := Dog{Name: "Kuckuc"}
	cat := Cat{Name: "Cata"}

	AnimalBehavior(dog)
	AnimalBehavior(cat)
}
