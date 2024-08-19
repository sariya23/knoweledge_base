package main

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct {
	Breed string
}

func (d Dog) Sound() {
	fmt.Println("Гав")
}

func (d Dog) String() string {
	return "Песель с породой " + d.Breed
}

type Cat struct {
	Breed string
}

func (c Cat) Sound() {
	fmt.Println("Мяу")
}

func WhatAnimal(a Animal) {
	switch a.(type) {
	case Cat:
		cat := a.(Cat)
		fmt.Println(cat.Breed)
	case Dog:
		dog := a.(Dog)
		fmt.Println(dog.Breed)
	default:
		fmt.Println("что-то непонятное")
	}
}

func main() {
	dog := Dog{"Овчарка"}
	fmt.Println(dog)
}
